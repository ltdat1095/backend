package routers

import (
	"backend/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/users", &controllers.UserController{}, "get:ListUsers")
	beego.Router("/signup", &controllers.UserController{}, "get:NewUserGet;post:NewUserPost")
}
