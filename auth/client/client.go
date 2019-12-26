package client

import (
	"context"
	"fmt"
	authmodel "outback/micro-go/auth/model"
	"outback/micro-go/auth/model/access"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
)

func MakeAccessToken(r *authmodel.Request) (string, error) {
	service := micro.NewService()
	service.Init()
	c := service.Client()
	// 请求参数
	req := access.Subject{
		ID:   "hello",
		Name: r.UserName,
	}
	request := c.NewRequest("mu.micro.book.srv.auth", "AuthService.MakeAccessToken", req, client.WithContentType("application/json"))
	var token string
	if err := c.Call(context.TODO(), request, &token); err != nil {
		fmt.Println(err)
		return "", err
	}
	return token, nil
}
