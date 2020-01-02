package client

import (
	"context"
	"outback/micro-go/api/entity"
	"outback/micro-go/api/service"

	"github.com/opentracing/opentracing-go"

	"github.com/micro/go-micro/util/log"

	hystrix_go "github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
)

var (
	userClient  client.Client
	span        opentracing.Span
	callWrapper client.CallWrapper
)

func Init() {
	hystrix_go.DefaultVolumeThreshold = 2
	hystrix_go.DefaultErrorPercentThreshold = 50
	hystrix_go.DefaultTimeout = 1000 * 4
	//userconfig := hystrix_go.CommandConfig{Timeout: hystrix_go.DefaultTimeout / 2}
	//hystrix_go.ConfigureCommand("GET-/user/login", userconfig)
	//userClient = breaker.NewUserClientWrapper()(client.DefaultClient)
	//userClient = client.DefaultClient
	//t, io, err := tracer.NewTracer("user-web", "")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer io.Close()
	//userClient = microTraceing.NewClientWrapper(t)(client.DefaultClient)
	//span = t.StartSpan("say-hello")
	//span.SetTag("hello-to", "UserClient")
	//callWrapper = microTraceing.NewCallWrapper(t)
	//defer span.Finish()

	//userClient = client.NewClient(
	//	// set the selector
	//	// add the trace wrapper
	//	client.Wrap(microTraceing.NewClientWrapper(t)),
	//)

	//userClient.Init(
	//	client.Retries(3),
	//	为了调试看log方便，始终返回true, nil，即会一直重试直至重试次数用尽
	//client.Retry(func(ctx context.Context, req client.Request, retryCount int, err error) (bool, error) {
	//	log.Log(req.Method(), retryCount, " client retry")
	//	return true, nil
	//}),
	//)
}

func QueryUserByName(ctx context.Context, name string) (*entity.User, error) {
	userService := service.NewUserService(userClient)
	request := userService.Clint.NewRequest(userService.Name, "UserHandler.QueryUserByName", name, client.WithContentType("application/json"))
	//request := userClient.NewRequest(userService.Name, "UserHandler.QueryUserByName", name, client.WithContentType("application/json"))
	response := new(entity.User)
	log.Info("client 开始调用服务")
	err := userService.Clint.Call(ctx, request, response)
	if err != nil {
		log.Infof("服务调用出错了 error is %s\n", err.Error())
		return response, err
	}
	return response, nil
}

func GetToken(ctx context.Context, user *entity.User) (string, error) {
	service := micro.NewService()
	service.Init()
	c := service.Client()
	// 请求参数

	request := c.NewRequest("mu.micro.book.srv.auth", "Service.MakeAccessToken", user, client.WithContentType("application/json"))
	response := new(entity.User)
	if err := c.Call(ctx, request, response); err != nil {
		log.Info("get token err", err)
		return "", err
	}
	return response.Token, nil
}
