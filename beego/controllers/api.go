package controllers

import (
	"encoding/json"
	"errors"

	"github.com/astaxie/beego"
	"github.com/ltdat/backend/logic"
)

type ApiController struct {
	beego.Controller
}

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *ApiController) Login() {
	var loginParams LoginParams
	json.Unmarshal(c.Ctx.Input.RequestBody, &loginParams)

	if loginParams.Email == "" || loginParams.Password == "" {
		c.responseFailed(400, errors.New("Bad request"))
		return
	}

	userLogic := logic.UserLogic{}
	user, err := userLogic.AuthUser(loginParams.Email, loginParams.Password)
	if err != nil || user.Id < 1 {
		c.responseFailed(400, errors.New("Invalid email or password"))
		return
	}

	c.responseSucceed()
	return
}

func (c *ApiController) SignUp() {
	var signUpParams logic.SignUpParams
	json.Unmarshal(c.Ctx.Input.RequestBody, &signUpParams)

	userLogic := logic.UserLogic{}
	err := userLogic.NewUser(signUpParams)
	if err != nil {
		c.responseFailed(500, err)
		return
	}

	c.responseSucceed()
	return
}

func (c *ApiController) responseFailed(statusCode int, err error) {
	beego.Error(err.Error())
	c.Data["json"] = map[string]interface{}{
		"status":  statusCode,
		"message": err.Error(),
	}
	c.ServeJSON()
}

func (c *ApiController) responseSucceed() {
	c.Data["json"] = map[string]interface{}{
		"status":  200,
		"message": "OK",
	}
	c.ServeJSON()
}
