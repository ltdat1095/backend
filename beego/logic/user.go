package logic

import (
	"errors"
	"fmt"
	"time"

	"github.com/ltdat/backend/models"
)

type UserLogic struct {
}

func (u *UserLogic) AuthUser(email string, password string) (*models.User, error) {
	user := &models.User{
		Email: email,
	}

	if err := user.Read("Email"); err != nil {
		errorMess := fmt.Sprintf("Can't find %s user", email)
		return nil, errors.New(errorMess)
	}

	if !CheckPass(password, user.Password) {
		return nil, errors.New("Invalid password")
	}

	user.LastLoginTime = time.Now()
	user.Update("LastLoginTime")
	return user, nil
}

type SignUpParams struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

func (u *UserLogic) NewUser(params SignUpParams) error {
	if models.Users().Filter("email", params.Email).Exist() {
		return errors.New("User already registered")
	}

	hashedPass, err := HashPass(params.Password)
	if err != nil {
		return errors.New("Failed to hash password")
	}

	user := models.User{
		Email:     params.Email,
		Password:  hashedPass,
		Age:       params.Age,
		FirstName: params.FirstName,
		LastName:  params.LastName,
	}
	err = user.Insert()
	if err != nil {
		return errors.New("Failed to insert new user")
	}

	return nil
}
