package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	internal "github.com/netivitton/go-rest-api/internal/routers"
)

func main() {
	env := os.Getenv("APP_ENV")
	fmt.Println("ENV:", env)
	godotenv.Load(env + ".env")
	routersInit := internal.InitRouter()
	server := &http.Server{
		Addr:    ":8080",
		Handler: routersInit,
	}
	server.ListenAndServe()
}
