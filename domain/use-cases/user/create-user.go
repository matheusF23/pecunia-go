package usecases

import "github.com/matheusF23/pecunia-go/domain/entity"

type CreateUser interface {
	Execute(user *entity.User) error
}
