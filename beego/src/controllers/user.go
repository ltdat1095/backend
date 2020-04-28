package controllers

import (
	"backend/logic"
	"backend/models"

	"github.com/astaxie/beego"
)

type UserController struct {
	BaseController
}

func (c *UserController) ListUsers() {
	var users []*models.User
	var err error

	if users, err = models.GetUsers(); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	c.Data["json"] = &users
	c.ServeJSON()
}

func (c *UserController) NewUserGet() {
	c.activeContent("user/new")
}

func (c *UserController) NewUserPost() {
	flash := beego.NewFlash()
	userLogic := logic.UserLogic{}

	user := models.User{}
	if err := c.ParseForm(&user); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	if !user.IsEmpty() {
		user, err := userLogic.NewUser(user)
		if err != nil || user.ID < 1 {
			flash.Error(err.Error())
			flash.Store(&c.Controller)
			return
		}
		flash.Success("New user createad")
		flash.Store(&c.Controller)
		return
	}
}
