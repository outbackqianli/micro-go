package main

import (
	"outback/micro-go/api/constent"
	"outback/micro-go/basic"
	"outback/micro-go/basic/db"
	"outback/micro-go/user-srv/handler"

	"github.com/micro/go-micro/util/log"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

func main() {
	// 初始化配置、数据库等信息
	basic.Init()

	// 新建服务
	service := micro.NewService(
		micro.Name(constent.ServiceName),
		micro.Version("latest"),
	)

	// 服务初始化
	service.Init(micro.Action(func(c *cli.Context) {
		db.Init()
	}))
	if err := micro.RegisterHandler(service.Server(), new(handler.UserHandler)); err != nil {
		log.Fatal(err)
	}
	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
