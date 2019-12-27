package main

import (
	"log"
	"outback/micro-go/basic"
	"outback/micro-go/user-web/handler"

	"github.com/gorilla/mux"

	"github.com/micro/go-micro/registry"

	"github.com/micro/cli"
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
				//handler.Init()
			}),
	); err != nil {
		log.Fatal(err)
	}

	// 注册登录接口

	//r := mux.NewRouter().Path("/user/login").Methods("GET").HandlerFunc(handler.Login)
	//r := mux.NewRouter()
	//http.Handle("/", r)
	//micro.RegisterHandler()

	r := mux.NewRouter()
	r.Path("/user/login").Methods("GET").HandlerFunc(handler.Login)

	service.Handle("/", r)

	//service.HandleFunc("/user/login", handler.Login)
	//service.HandleFunc("/user/login2", handler.Login)
	//service.HandleFunc("/user/login3", handler.Login)
	//service.HandleFunc("/user/login4", handler.Login)

	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
