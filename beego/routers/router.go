package routers

import (
	"github.com/astaxie/beego"
	"github.com/ltdat/backend/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/signup", &controllers.SignUpController{})

	beego.Router("/api/login", &controllers.ApiController{}, "post:Login")
	beego.Router("/api/signUp", &controllers.ApiController{}, "post:SignUp")
}
