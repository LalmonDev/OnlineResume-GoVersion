// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"OnlineResume-GoVersion/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	//ns := beego.NewNamespace("/v1",
	//	beego.NSNamespace("/object",
	//		beego.NSInclude(
	//			&controllers.ObjectController{},
	//		),
	//	),
	//	beego.NSNamespace("/user",
	//		beego.NSInclude(
	//			&controllers.UserController{},
	//		),
	//	),
	//
	//	// 自己实现的例子，练习用
	//	beego.NSNamespace("/Hello",
	//		beego.NSGet("h",func(ctx *context.Context) {
	//		ctx.Output.Body([]byte("Hello"))
	//	}),
	//	),
	//)
	//beego.AddNamespace(ns)

	// 路由注册
	beego.Router("/Hello",&controllers.HelloController{},"*:Hello")

}
