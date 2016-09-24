package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"quickstart/controllers"
)

func init() {
	beego.Get("/myrouter",func(ctx *context.Context){
		ctx.Output.Body([]byte("hello world myrouter"))
	})

	beego.Router("/deny",&controllers.RController{})
}
