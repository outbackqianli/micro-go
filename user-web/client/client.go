package client

import (
	"context"
	"fmt"
	"outback/micro-go/user-srv/model"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
)

func QueryUserByName(name string) (*model.User, error) {
	service := micro.NewService()
	service.Init()
	c := service.Client()
	// 请求参数
	request := c.NewRequest("mu.micro.book.srv.user", "User.QueryUserByName", name, client.WithContentType("application/json"))
	response := new(model.User)
	if err := c.Call(context.TODO(), request, response); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return response, nil
}

func GetToken(user *model.User) (string, error) {
	service := micro.NewService()
	service.Init()
	c := service.Client()
	// 请求参数

	request := c.NewRequest("mu.micro.book.srv.auth", "Service.MakeAccessToken", user, client.WithContentType("application/json"))
	response := new(model.User)
	if err := c.Call(context.TODO(), request, response); err != nil {
		fmt.Println(err)
		return "", err
	}
	return response.Token, nil
}
