package main

import (
	"outback/micro-go/basic"
	"outback/micro-go/plugins/breaker"
	userClient "outback/micro-go/user-web/client"
	"outback/micro-go/user-web/handler"

	"github.com/micro/go-micro/util/log"

	"github.com/gorilla/mux"

	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
)

func main() {
	// 初始化配置
	basic.Init()

	// 使用etcd注册
	//micReg := etcd.NewRegistry(registryOptions)
	//reg := memory.NewRegistry()
	reg := registry.DefaultRegistry
	// 创建新服务
	service := web.NewService(
		// 后面两个web，第一个是指是web类型的服务，第二个是服务自身的名字
		web.Name("mu.micro.book.web.user"),
		web.Version("latest"),
		web.Registry(reg),
		web.Address(":8088"),
	)

	// 初始化服务
	if err := service.Init(
		web.Action(
			func(c *cli.Context) {
				// 初始化handler
				userClient.Init()
			}),
	); err != nil {
		log.Fatal(err)
	}
	log.Debug("debug  ")
	log.Info("INfo")
	r := mux.NewRouter()
	// queries 表示必传参数，且只能成对出现
	r.Path("/user/login").Methods("GET").HandlerFunc(handler.Login)
	//hand := breaker.BreakerWrapper(handler.Login)

	//service.HandleFunc("/user/login", handler.Login)

	service.Handle("/", breaker.BreakerWrapper(r))
	//service.Handle("/", r)
	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
