package routers

import (
	"rgxtest/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/helloworld", &controllers.MainController{}, "get:HelloWorld")
}
