package routers

import (
	"OnlineResume-GoVersion/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
