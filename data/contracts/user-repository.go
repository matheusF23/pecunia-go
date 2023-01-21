package contracts

import (
	"github.com/matheusF23/pecunia-go/domain/entity"
)

type UserRepository interface {
	CreateUser(user *entity.User) error
}
