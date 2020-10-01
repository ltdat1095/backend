package controllers

import (
	"github.com/astaxie/beego"
)

const (
	sessName = "userLogin"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
	session := c.StartSession()
	userID := session.Get("UserID")
	if userID == nil {
		c.Redirect("/login", 302)
	}
}

// func (c *BaseController) Auth() {
// 	c.IsLogin = c.GetSession(sessName) != nil
// 	if !c.IsLogin {
// 		c.Redirect(c.LoginPath(), 302)
// 	}
// }

// func (c *BaseController) GetLoginUser() *models.User {
// 	u := &models.User{Id: c.GetSession(sessName).(int64)}
// 	u.Read()
// 	return u
// }

// func (c *BaseController) DelLogin() {
// 	c.DelSession(sessName)
// }

// func (c *BaseController) LoginPath() string {
// 	return c.UrlFor("LoginController.Login")
// }

func (c *BaseController) activeContent(view string) {
	c.Layout = "layouts/layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "layouts/scripts.html"
	c.TplName = view + ".html"
}
