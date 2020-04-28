package routers

import (
	"github.com/astaxie/beego"
	"github.com/ltdat/backend/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/users", &controllers.UserController{}, "get:ListUsers")
	beego.Router("/signup", &controllers.UserController{}, "get:NewUserGet;post:NewUserPost")
}
