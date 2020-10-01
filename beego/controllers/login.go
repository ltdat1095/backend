package controllers

type LoginController struct {
	BaseController
}

func (c *LoginController) Get() {
	c.activeContent("login/login")
}
