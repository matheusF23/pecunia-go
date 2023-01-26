package services

import (
	"github.com/matheusF23/pecunia-go/internal/data/contracts"
	"github.com/matheusF23/pecunia-go/internal/domain/entity"
)

type CreateUserService struct {
	userRepository contracts.UserRepository
}

func NewCreateUserService(ur contracts.UserRepository) *CreateUserService {
	return &CreateUserService{userRepository: ur}
}

func (createUserService *CreateUserService) ExecuteCreateUser(user *entity.User) error {
	err := createUserService.userRepository.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}
