package main

import (
	"github.com/netivitton/go-rest-api/routers"
)

func main() {
	r := routers.InitRouter()
	r.run()
}
