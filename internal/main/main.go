package main

import (
	"fmt"
	"net/http"

	"github.com/matheusF23/pecunia-go/internal/main/configs"
	"github.com/matheusF23/pecunia-go/internal/main/routes"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}
	router := routes.SetupRoutes()
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), router)
}
