package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/matheusF23/pecunia-go/internal/domain/entity"
	usecases "github.com/matheusF23/pecunia-go/internal/domain/use-cases/user"
)

type CreateUserController struct {
	CreateUser usecases.CreateUser
}

func NewCreateUserController(createUser usecases.CreateUser) *CreateUserController {
	return &CreateUserController{CreateUser: createUser}
}

func (c *CreateUserController) Handle(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = c.CreateUser.ExecuteCreateUser(&user)
	var response map[string]any
	if err != nil {
		response = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
	} else {
		response = map[string]any{
			"Error":   false,
			"Message": "User created successfully",
		}
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
