package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyID_WhenCreateANewUser_ThenShouldReceiveAnError(t *testing.T) {
	user := User{}
	assert.Error(t, user.IsValid(), "invalid id")
}

func TestGivenAnEmptyName_WhenCreateANewUser_ThenShouldReceiveAnError(t *testing.T) {
	user := User{ID: "123"}
	assert.Error(t, user.IsValid(), "invalid name")
}

func TestGivenAnEmptyEmail_WhenCreateANewUser_ThenShouldReceiveAnError(t *testing.T) {
	user := User{ID: "123", Name: "Bartolomeu"}
	assert.Error(t, user.IsValid(), "invalid email")
}

func TestGivenAnEmptyPassword_WhenCreateANewUser_ThenShouldReceiveAnError(t *testing.T) {
	user := User{ID: "123", Name: "Bartolomeu", Email: "barto@gmail.com"}
	assert.Error(t, user.IsValid(), "invalid password")
}

func TestGivenAValidParams_WhenCreateANewUser_ThenIShouldReceiveCreatedUserWithAllParams(t *testing.T) {
	user := User{ID: "123", Name: "Bartolomeu", Email: "barto@gmail.com", Password: "secretPass"}
	assert.Equal(t, "123", user.ID)
	assert.Equal(t, "Bartolomeu", user.Name)
	assert.Equal(t, "barto@gmail.com", user.Email)
	assert.Equal(t, "secretPass", user.Password)
	assert.Nil(t, user.IsValid())
}

func TestGivenAValidParams_WhenICallANewUserFunc_ThenIShouldReceiveCreatedUserWithAllParams(t *testing.T) {
	user, err := NewUser("123", "Bartolomeu", "barto@gmail.com", "secretPass")
	assert.Nil(t, err)
	assert.Equal(t, "123", user.ID)
	assert.Equal(t, "Bartolomeu", user.Name)
	assert.Equal(t, "barto@gmail.com", user.Email)
	assert.Equal(t, "secretPass", user.Password)
}
