package routers

import (
	"github.com/astaxie/beego"
	"quickstart/controllers"
)

func init() {
	beego.Router("/sample",&controllers.TragController{},"*:SampleDrag")
	beego.Router("/shopping",&controllers.TragController{},"*:Shopping")

}
