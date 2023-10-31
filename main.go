package main

import (
	"fmt"
	"log"
	"net/http"
	"pontos_funcionario/src/config"
	app_routes "pontos_funcionario/src/main/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	} else {
		fmt.Println("database configs loaded")
	}

	hostConfig, err := config.GetHostConfig()
	if err != nil {
		log.Panic(err)
	} else {
		fmt.Println("server configs loaded")
	}

	r := app_routes.MakeRoutes()
	http.ListenAndServe(fmt.Sprintf("%s:%d", hostConfig.Host, hostConfig.Port), r)
}
