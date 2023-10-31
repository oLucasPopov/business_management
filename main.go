package main

import (
	"fmt"
	"log"
	"net/http"
	"pontos_funcionario/src/config"
	app_routes "pontos_funcionario/src/main/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func startServer(r *mux.Router) (string, *mux.Router) {
	hostConfig, err := config.GetHostConfig()

	if err != nil {
		log.Panic(err)
	} else {
		fmt.Println("server configs loaded")
	}

	serverString := fmt.Sprintf("%s:%d", hostConfig.Host, hostConfig.Port)
	fmt.Println("Starting server at", serverString)

	return serverString, r
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	} else {
		fmt.Println("database configs loaded")
	}

	r := app_routes.MakeRoutes()
	http.ListenAndServe(startServer(r))
}
