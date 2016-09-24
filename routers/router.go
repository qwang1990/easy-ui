package routers

import (
	"quickstart/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/api/list", &controllers.TestController{}, "*:List")
	beego.Router("/person/:last/:first", &controllers.TestController{})
	beego.AutoRouter(&controllers.TestController{})

	ns :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/shop",
				beego.NSGet("/:id", func(ctx *context.Context) {
					ctx.Output.Body([]byte(ctx.Request.Host))
				}),
			),
			beego.NSNamespace("/order",
				beego.NSGet("/:id", func(ctx *context.Context) {
					ctx.Output.Body([]byte("orderinfo"))
				}),
			),
			beego.NSNamespace("/crm",
				beego.NSGet("/:id", func(ctx *context.Context) {
					ctx.Output.Body([]byte("crminfo"))
				}),
			),
		)

	beego.AddNamespace(ns)

}

