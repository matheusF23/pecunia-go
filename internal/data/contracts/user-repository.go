package contracts

import (
	"github.com/matheusF23/pecunia-go/internal/domain/entity"
)

type UserRepository interface {
	CreateUser(user *entity.User) error
}
