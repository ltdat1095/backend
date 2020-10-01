package controllers

type SignUpController struct {
	BaseController
}

func (c *SignUpController) Get() {
	c.activeContent("signUp/signUp")
}
