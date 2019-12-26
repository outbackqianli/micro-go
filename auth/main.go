package main

import (
	"log"
	"outback/micro-go/auth/handler"
	"outback/micro-go/auth/model"
	"outback/micro-go/basic"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

func main() {
	// 初始化配置、数据库等信息
	basic.Init()

	// 新建服务
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.auth"),
		micro.Version("latest"),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化handler
			model.Init()
			// 初始化handler
			handler.Init()
		}),
	)

	// 注册服务
	micro.RegisterHandler(service.Server(), new(handler.AuthService))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
