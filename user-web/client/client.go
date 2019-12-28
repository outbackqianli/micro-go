package client

import (
	"context"
	"fmt"
	"outback/micro-go/api/entity"
	"outback/micro-go/api/service"

	"github.com/micro/go-micro/util/log"

	hystrix_go "github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-plugins/wrapper/breaker/hystrix"

	"github.com/micro/go-micro"
)

var (
	userClient client.Client
)

func Init() {
	hystrix_go.DefaultVolumeThreshold = 1
	hystrix_go.DefaultErrorPercentThreshold = 1
	userClient = hystrix.NewClientWrapper()(client.DefaultClient)
	userClient.Init(
		client.Retries(3),
		//为了调试看log方便，始终返回true, nil，即会一直重试直至重试次数用尽
		client.Retry(func(ctx context.Context, req client.Request, retryCount int, err error) (bool, error) {
			log.Log(req.Method(), retryCount, " client retry")
			return true, nil
		}),
	)
}

func QueryUserByName(name string) (*entity.User, error) {
	fmt.Println("web is there QueryUserByName")
	userService := service.NewUserService(userClient)
	request := userService.Clint.NewRequest(userService.Name, "UserHandler.QueryUserByName", name, client.WithContentType("application/json"))
	response := new(entity.User)
	if err := userService.Clint.Call(context.TODO(), request, response); err != nil {
		fmt.Println(err)
		return nil, err
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
