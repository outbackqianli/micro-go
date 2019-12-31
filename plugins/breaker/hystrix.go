/*
这个包可以进行服务的熔断
*/

package breaker

import (
	"context"
	"outback/micro-go/api/entity"

	"github.com/micro/go-micro/util/log"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/client"
)

type userClientWrapper struct {
	client.Client
}

func (c *userClientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	runFunc := func() error {
		log.Info("熔断执行开始runFunc")
		err := c.Client.Call(ctx, req, rsp, opts...)
		log.Infof("熔断执行完Call 方法 error is %s \n", err)
		return err
	}
	fallbackFunc := func(err error) error {
		if err != nil {
			log.Infof("熔断执行fallbackFunc error is ", err.Error())
			rspUser, _ := rsp.(*entity.User)
			rspUser.Id = 12
			rspUser.Name = "熔断之后的返回"
			rspUser.Pwd = "熔断之后的返回"
			return nil
		}
		return err
	}
	commandName := req.Service() + "." + req.Endpoint()
	log.Info("熔断name is ", commandName)
	//output := make(chan bool, 1)

	errs := hystrix.Do("GET-/user/login", runFunc, fallbackFunc)
	//select {
	//case _ = <-output:
	//	return nil
	//case err := <-errs:
	//	return err
	//}
	return errs
}

// NewUserClientWrapper returns a hystrix client Wrapper.
func NewUserClientWrapper() client.Wrapper {
	fn := func(c client.Client) client.Client {
		return &userClientWrapper{c}
	}
	return fn
}
