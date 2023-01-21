package services

import (
	"github.com/matheusF23/pecunia-go/data/contracts"
	"github.com/matheusF23/pecunia-go/data/models"
)

type CreateUserService struct {
	userRepository contracts.UserRepository
}

func (createUserService *CreateUserService) Execute(user *models.UserModel) error {
	err := createUserService.userRepository.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}
