package main

import (
	"outback/micro-go/basic"
	"outback/micro-go/basic/db"
	"outback/micro-go/user-srv/model"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/sirupsen/logrus"
)

func main() {
	// 初始化配置、数据库等信息
	basic.Init()

	// 新建服务
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.user"),
		micro.Version("latest"),
	)

	// 服务初始化
	service.Init(micro.Action(func(c *cli.Context) {
		db.Init()
	}))
	micro.RegisterHandler(service.Server(), new(model.User))

	// 启动服务
	if err := service.Run(); err != nil {
		logrus.Fatal(err)
	}
}
