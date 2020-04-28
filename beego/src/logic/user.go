package logic

import (
	"errors"
	"fmt"
	"backend/models"
	"time"

	"github.com/astaxie/beego"
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

func (u *UserLogic) NewUser(user models.User) (*models.User, error) {
	if models.Users().Filter("email", user.Email).Exist() {
		return nil, errors.New("User already registered")
	}

	var (
		hashedPass string
		err        error
	)

	if hashedPass, err = HashPass(user.Password); err != nil {
		return nil, errors.New("Failed to hash password")
	}
	user.Password = hashedPass

	err = user.Insert()
	if err != nil {
		beego.Debug(err)
		return nil, errors.New("Failed to insert new user")
	}

	return &user, err
}
