package main

import (
	"fmt"

	"github.com/micro/go-micro/util/log"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/file"
)

func main() {
	//conf := config.NewConfig()
	//conf.Load(
	//	file.NewSource(
	//		file.WithPath(""))
	//
	//	)
	// load the config from a file source
	if err := config.Load(file.NewSource(
		file.WithPath("/Users/liuqiang/workspace/golang/src/outback/micro-go/configration/config-dev.json"),
	)); err != nil {
		fmt.Println(err)
		return
	}
	//c := config.Map()
	//log.Info("map is ", c)
	// define our own host type
	type Host struct {
		Address string `json:"address"`
		Port    int    `json:"port"`
	}
	address := config.Get("hosts", "database", "address").String("localhost")
	log.Info("address is ", address)
	w, err := config.Watch("hosts", "database")
	if err != nil {
		log.Info("watch error ", err.Error())
		return
	}

	// wait for next value
	v, err := w.Next()
	if err != nil {
		// do something
		log.Info("next error ", err.Error())
		return
	}

	var host Host

	err = v.Scan(&host)
	if err != nil {
		log.Info("scran err ", err.Error())
		return
	}
	//var host Host

	// read a database host
	//if err := config.Get("hosts", "database").Scan(&host); err != nil {
	//	fmt.Println(err)
	//	return
	//}

	fmt.Println("address and host ", host.Address, host.Port)
	//err := config.LoadFile("/Users/liuqiang/workspace/golang/src/outback/micro-go/configration/config-dev.json")
	//log.Info("conf", err)

}

type DataBase struct {
	Host string `json:"host"`
	Port string `json:"port"`
}
