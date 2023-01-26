package entity

import "errors"

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(id string, name string, email string, password string) (*User, error) {
	user := &User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}
	err := user.IsValid()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (user *User) IsValid() error {
	if user.ID == "" {
		return errors.New("invalid id")
	}
	if user.Name == "" {
		return errors.New("invalid name")
	}
	if user.Email == "" {
		return errors.New("invalid email")
	}
	if user.Password == "" {
		return errors.New("invalid password")
	}
	return nil
}
