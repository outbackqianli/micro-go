package main

import (
	"log"
	"time"

	"github.com/micro/go-micro"

	"context"
)

type Say struct{}

type Requst struct {
	Name string `json:"name"`
}
type Response struct {
	Name string `json:"name"`
	Msg  string `json:"msg"`
}

func (s *Say) HelloWord(ctx context.Context, req *Requst, rsp *Response) error {
	log.Print("Received Say.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.greeter"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	micro.RegisterHandler(service.Server(), new(Say))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
