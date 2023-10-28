package main

import (
	"fmt"
	"log"
	"net/http"
	app_routes "pontos_funcionario/src/main/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	} else {
		fmt.Print(".env loaded")
	}

	r := app_routes.MakeRoutes()
	http.ListenAndServe("localhost:8080", r)
}
