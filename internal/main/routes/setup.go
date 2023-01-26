package routes

import (
	"github.com/go-chi/chi/v5"
)

func SetupRoutes() chi.Router {
	router := chi.NewRouter()
	router.Mount("/users", UserRoute())
	return router
}
