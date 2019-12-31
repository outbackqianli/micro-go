package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"outback/micro-go/basic"

	"github.com/gorilla/mux"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/prometheus/common/log"
)

func main() {
	// 初始化配置
	basic.Init()

	reg := registry.DefaultRegistry
	// 创建新服务
	service := web.NewService(
		// 后面两个web，第一个是指是web类型的服务，第二个是服务自身的名字
		web.Name("mu.micro.book.web.user"),
		web.Version("latest"),
		web.Registry(reg),
		web.Address(":10010"),
	)

	// 初始化服务
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter().StrictSlash(true)
	// queries 表示必传参数，且只能成对出现
	r.Path("/config/mysql").Methods("GET").HandlerFunc(MysqlConfig)
	service.Handle("/", r)

	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

// Login 登录入口
func MysqlConfig(w http.ResponseWriter, r *http.Request) {
	log.Info("请求一次")
	type Database struct {
		Name   string `json:"name"`
		Ticker int    `json:"ticker"`
	}
	w.WriteHeader(http.StatusOK)
	d := Database{Name: fmt.Sprintf("mysql +%d", 100), Ticker: 2}
	u, _ := json.Marshal(d)
	w.Write(u)
}
