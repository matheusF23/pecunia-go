package services

import (
	"github.com/matheusF23/pecunia-go/data/contracts"
	"github.com/matheusF23/pecunia-go/domain/entity"
)

type CreateUserService struct {
	userRepository contracts.UserRepository
}

func (createUserService *CreateUserService) Execute(user *entity.User) error {
	err := createUserService.userRepository.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}
