package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/matheusF23/pecunia-go/internal/main/factories"
)

func UserRoute() chi.Router {
	userRouter := chi.NewRouter()
	userRouter.Post("/", factories.MakeCreateUserController().Handle)
	return userRouter
}
