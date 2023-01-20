package contracts

import "github.com/matheusF23/pecunia-go/data/models"

type UserRepository interface {
	CreateUser(user models.UserModel) error
}
