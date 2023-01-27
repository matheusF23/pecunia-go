package database

import (
	"database/sql"

	"github.com/matheusF23/pecunia-go/internal/domain/entity"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (userRepository *UserRepository) CreateUser(user *entity.User) error {
	stmt, err := userRepository.Db.Prepare("INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}
