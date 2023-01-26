package usecases

import "github.com/matheusF23/pecunia-go/internal/domain/entity"

type CreateUser interface {
	ExecuteCreateUser(user *entity.User) error
}
