package main

import (
	"fmt"
	"net"
	"net/http"
	"outback/micro-go/basic"
	"outback/micro-go/basic/config"
	tracer "outback/micro-go/plugins/tracer/jaeger"
	"outback/micro-go/plugins/tracer/opentracing/std2micro"
	userClient "outback/micro-go/user-web/client"
	"outback/micro-go/user-web/client/breaker"
	"outback/micro-go/user-web/handler"
	"time"

	"github.com/micro/go-plugins/registry/kubernetes"

	"github.com/afex/hystrix-go/hystrix"

	"github.com/opentracing/opentracing-go"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"

	"github.com/gorilla/mux"

	"github.com/micro/cli"
	"github.com/micro/go-micro/web"
)

func main() {
	// 初始化配置
	basic.Init()
	// 使用etcd注册
	//micReg := etcd.NewRegistry(registryOptions)
	//reg := memory.NewRegistry()

	t, io, err := tracer.NewTracer("user-web-url", "")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	reg := kubernetes.NewRegistry(registryOptions)
	// 创建新服务
	service := web.NewService(
		// 后面两个web，第一个是指是web类型的服务，第二个是服务自身的名字
		web.Name("mu.micro.book.web.user"),
		web.Version("latest"),
		web.Registry(reg),
		web.Address(":8088"),
		web.RegisterTTL(time.Second*15),
		web.RegisterInterval(time.Second*10),
	)
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(net.JoinHostPort("", "81"), hystrixStreamHandler)

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
	//设置采样率
	std2micro.SetSamplingFrequency(100)

	log.Debug("debug  ")
	log.Info("INfo")
	r := mux.NewRouter()
	// queries 表示必传参数，且只能成对出现
	r.Path("/user/login").Methods("GET").HandlerFunc(handler.Login)
	//r.Path("/user/logintwo").Methods("GET").HandlerFunc(handler.Login2)
	//service.Handle("/", breaker.BreakerWrapper(r))
	//增加链路追踪
	service.Handle("/", std2micro.TracerWrapper(breaker.BreakerWrapper(r)))
	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
