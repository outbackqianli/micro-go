package breaker

import (
	"context"
	"errors"
	"fmt"

	"github.com/micro/go-micro/util/log"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/client"
)

type clientWrapper struct {
	client.Client
}

func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	fmt.Println("befor res", rsp, "\n")
	runFunc := func() error {
		log.Info("服务熔断开始时 run Func")
		err := c.Client.Call(ctx, req, rsp, opts...)
		log.Info("服务熔断开始时 run Func error is ", err)
		return err
	}
	fallbackFunc := func(err error) error {
		if err != nil {
			log.Info("服务熔断开始了,执行fallbackFunc")
		}
		return errors.New("fuck")
	}

	commandName := req.Service() + "." + req.Endpoint()
	err := hystrix.Do(commandName, runFunc, fallbackFunc)
	fmt.Printf("after res is  %+v,the error is %s \n", rsp, err)

	return err
}

// NewClientWrapper returns a hystrix client Wrapper.
func NewClientWrapper() client.Wrapper {
	fn := func(c client.Client) client.Client {
		return &clientWrapper{c}
	}
	return fn
}
