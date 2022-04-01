package main

import (
	"net/http"

	"github.com/netivitton/go-rest-api/routers"
)

func main() {
	routersInit := routers.InitRouter()
	server := &http.Server{
		Addr:    ":8080",
		Handler: routersInit,
	}
	server.ListenAndServe()
}
