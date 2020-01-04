/*
这个包可以进行服务的熔断
*/

package breaker

import (
	"context"
	"net/http"
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
		//if err == hystrix.ErrCircuitOpen {
		//	log.Infof("降级fallbackFunc error is ", err.Error())
		//	rspUser, _ := rsp.(*entity.User)
		//	rspUser.Id = 13
		//	rspUser.Name = "降级之后的返回"
		//	rspUser.Pwd = "降级之后的返回"
		//	return nil
		//}
		//if err != hystrix.ErrCircuitOpen && err != nil {
		log.Infof("熔断执行fallbackFunc error is ", err.Error())
		rspUser, _ := rsp.(*entity.User)
		rspUser.Id = 12
		rspUser.Name = "熔断之后的返回"
		rspUser.Pwd = "熔断之后的返回"
		return nil
		//}
		//return err
	}
	commandName := req.Service() + "." + req.Endpoint()
	log.Info("熔断name is ", commandName)

	errs := hystrix.Do(commandName, runFunc, fallbackFunc)
	return errs
}

// NewUserClientWrapper returns a hystrix client Wrapper.
func NewUserClientWrapper() client.Wrapper {
	fn := func(c client.Client) client.Client {
		return &userClientWrapper{c}
	}
	return fn
}

//BreakerWrapper hystrix breaker
func BreakerWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.Method + "-" + r.RequestURI
		runFunc := func() error {
			h.ServeHTTP(w, r)
			return nil
		}
		fallFunc := func(err error) error {
			if err == hystrix.ErrCircuitOpen {
				w.WriteHeader(http.StatusAccepted)
				w.Write([]byte("请稍后重试"))
				return nil
			}
			h.ServeHTTP(w, r)
			return err
		}
		hystrix.Do(name, runFunc, fallFunc)
	})
}
