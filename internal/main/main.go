package main

import (
	"fmt"
	"net/http"

	"github.com/matheusF23/pecunia-go/internal/main/configs"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}
	router := configs.SetupRoutes()
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), router)
}
