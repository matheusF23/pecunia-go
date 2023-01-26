package configs

import (
	"github.com/go-chi/chi/v5"
	"github.com/matheusF23/pecunia-go/internal/main/routes"
)

func SetupRoutes() chi.Router {
	router := chi.NewRouter()
	router.Mount("/users", routes.User())
	return router
}
