package main

import (
	"net/http"
	app_routes "pontos_funcionario/src/main/routes"
)

func main() {
	r := app_routes.MakeRoutes()
	http.ListenAndServe("localhost:8080", r)
}
