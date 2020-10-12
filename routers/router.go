package routers

import (
	"Data_Attestatio/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/user_register", &controllers.RegisterController{})

    beego.Router("/login.html", &controllers.LoginControllers{})

    beego.Router("/user_login", &controllers.LoginControllers{})
}
