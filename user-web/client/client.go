package client

import (
	"context"
	"fmt"
	"outback/micro-go/api/entity"
	"outback/micro-go/api/service"
	"outback/micro-go/plugins/breaker"

	hystrix_go "github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
)

var (
	userClient client.Client
)

func Init() {
	hystrix_go.DefaultVolumeThreshold = 1
	hystrix_go.DefaultErrorPercentThreshold = 1
	hystrix_go.DefaultTimeout = 1000 * 1
	userClient = breaker.NewUserClientWrapper()(client.DefaultClient)
	//userClient.Init(
	//	client.Retries(3),
	//	为了调试看log方便，始终返回true, nil，即会一直重试直至重试次数用尽
	//client.Retry(func(ctx context.Context, req client.Request, retryCount int, err error) (bool, error) {
	//	log.Log(req.Method(), retryCount, " client retry")
	//	return true, nil
	//}),
	//)
}

func QueryUserByName(name string) (*entity.User, error) {
	userService := service.NewUserService(userClient)
	request := userService.Clint.NewRequest(userService.Name, "UserHandler.QueryUserByName", name, client.WithContentType("application/json"))
	response := new(entity.User)
	fmt.Println("client 开始调用服务")
	err := userService.Clint.Call(context.TODO(), request, response)
	if err != nil {
		fmt.Printf("服务调用出错了 error is %s\n", err.Error())
		//fmt.Printf("此时response is %+v \n", response)
		return response, err
	}
	return response, nil
}

func GetToken(user *entity.User) (string, error) {
	service := micro.NewService()
	service.Init()
	c := service.Client()
	// 请求参数

	request := c.NewRequest("mu.micro.book.srv.auth", "Service.MakeAccessToken", user, client.WithContentType("application/json"))
	response := new(entity.User)
	if err := c.Call(context.TODO(), request, response); err != nil {
		fmt.Println(err)
		return "", err
	}
	return response.Token, nil
}
