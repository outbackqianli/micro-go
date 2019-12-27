package main

import (
	"log"

	"github.com/gorilla/mux"
	httpapi "github.com/micro/go-micro/api/server/http"
)

func main() {
	//var h http.Handler
	r := mux.NewRouter()
	api := httpapi.NewServer(":8089")
	//h = r
	api.Handle("/", r)
	// Start API
	if err := api.Start(); err != nil {
		log.Fatal(err)
	}
}
