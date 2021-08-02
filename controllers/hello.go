package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type HelloController struct {
	beego.Controller
}

func (ptr *HelloController) Hello() {
	logs.Debug("%+v", *ptr.Ctx.Input)
	hello := "Hello Beego Test!"
	resp := make(map[string]interface{})
	resp["Method"] = ptr.Ctx.Input.Method()
	resp["IP"] = ptr.Ctx.Input.IP()
	resp["Msg"] = hello

	ptr.Data["json"] = resp
	ptr.ServeJSON()
}
