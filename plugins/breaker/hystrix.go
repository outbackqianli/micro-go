/*
这个包可以进行服务的熔断
*/

package breaker

import (
	"context"
	"fmt"
	"outback/micro-go/api/entity"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/client"
)

type userClientWrapper struct {
	client.Client
}

func (c *userClientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	runFunc := func() error {
		fmt.Println("熔断执行开始runFunc")
		err := c.Client.Call(ctx, req, rsp, opts...)
		fmt.Printf("熔断执行完Call 方法 error is %s \n", err)
		return err
	}
	fallbackFunc := func(err error) error {
		if err != nil {
			fmt.Println("熔断执行fallbackFunc")
			rspUser, _ := rsp.(*entity.User)
			rspUser.Id = 12
			rspUser.Name = "熔断之后的返回"
			rspUser.Pwd = "熔断之后的返回"

		}
		return err
	}

	commandName := req.Service() + "." + req.Endpoint()
	err := hystrix.Do(commandName, runFunc, fallbackFunc)

	return err
}

// NewUserClientWrapper returns a hystrix client Wrapper.
func NewUserClientWrapper() client.Wrapper {
	fn := func(c client.Client) client.Client {
		return &userClientWrapper{c}
	}
	return fn
}
