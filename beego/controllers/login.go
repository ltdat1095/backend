package controllers

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.Layout = "layouts/layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "layouts/scripts.html"
	c.TplName = "login/login.html"
}
