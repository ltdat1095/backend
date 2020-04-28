package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/ltdat/backend/logic"
)

type LoginController struct {
	BaseController
}

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *LoginController) Get() {
	c.activeContent("login/login")
}

func (c *LoginController) Post() {
	var loginParams LoginParams
	json.Unmarshal(c.Ctx.Input.RequestBody, &loginParams)

	beego.Debug(loginParams)

	userLogic := logic.UserLogic{}
	if loginParams.Email != "" && loginParams.Password != "" {
		user, err := userLogic.AuthUser(loginParams.Email, loginParams.Password)
		if err != nil || user.ID < 1 {
			c.Data["json"] = map[string]interface{}{
				"status":  "400",
				"message": "Please try again",
			}
			c.ServeJSON()
		}
		mess := "Welcome " + user.Email
		c.Data["json"] = map[string]interface{}{
			"status":  "200",
			"message": mess,
		}
		c.ServeJSON()
	}
}

// func (c *LoginController) Logout() {
// 	c.DelLogin()
// 	flash := beego.NewFlash()
// 	flash.Success("Success logged out")
// 	flash.Store(&c.Controller)

// 	c.Ctx.Redirect(302, c.UrlFor("LoginController.Login"))
// }

// func (c *LoginController) Signup() {
// 	c.TplNames = "login/signup.tpl"
// 	c.Data["xsrfdata"] = template.HTML(c.XsrfFormHtml())

// 	if !c.Ctx.Input.IsPost() {
// 		return
// 	}

// 	var err error
// 	flash := beego.NewFlash()

// 	u := &models.User{}
// 	if err = c.ParseForm(u); err != nil {
// 		flash.Error("Signup invalid!")
// 		flash.Store(&c.Controller)
// 		return
// 	}
// 	if err = models.IsValid(u); err != nil {
// 		flash.Error(err.Error())
// 		flash.Store(&c.Controller)
// 		return
// 	}

// 	id, err := lib.SignupUser(u)
// 	if err != nil || id < 1 {
// 		flash.Warning(err.Error())
// 		flash.Store(&c.Controller)
// 		return
// 	}

// 	flash.Success("Register user")
// 	flash.Store(&c.Controller)

// 	c.SetLogin(u)

// 	c.Redirect(c.UrlFor("UsersController.Index"), 303)
// }
