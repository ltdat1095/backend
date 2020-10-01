package controllers

import "github.com/astaxie/beego"

type SignUpController struct {
	beego.Controller
}

func (c *SignUpController) Get() {
	c.Layout = "layouts/layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "layouts/scripts.html"
	c.TplName = "signUp/signUp.html"
}
