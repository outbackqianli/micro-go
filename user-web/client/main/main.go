package main

import (
	"fmt"

	"github.com/micro/go-micro/registry"

	"github.com/micro/go-micro/util/log"
)

func main() {
	//reg := memory.NewRegistry()
	reg := registry.DefaultRegistry
	getService, err := reg.GetService("mu.micro.book.web.user")
	if err != nil {
		log.Error("GetService error :", err.Error())
		return
	}
	for i := 0; i < len(getService); i++ {
		s := getService[i]
		fmt.Printf("s is %+v \n", s.Endpoints[0].Name)
		//fmt.Printf("s is %+v \n", s.Endpoints[1].Name)
		//fmt.Printf("s is %+v \n", s.Endpoints[2].Name)
		//fmt.Printf("s is %+v \n", s.Endpoints[3].Name)
		//fmt.Printf("s is %+v \n", s.Endpoints[1].Request)
		//fmt.Printf("s is %+v \n", s.Endpoints[2].Response)
		fmt.Printf("s is %+v \n", s.Nodes[0].Address)
		fmt.Printf("s is %+v \n", s.Nodes[0].Metadata)
	}
}
