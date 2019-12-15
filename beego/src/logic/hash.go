package logic

import (
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
)

func HashPass(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPass(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	beego.Debug(err)
	return err == nil
}
