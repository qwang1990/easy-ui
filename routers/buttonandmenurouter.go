package routers

import (
	"github.com/astaxie/beego"
	"quickstart/controllers"
)

func init() {
	beego.Router("/samplemenu",&controllers.ButtonController{},"*:SampleMenu")
	beego.Router("/linkbutton",&controllers.ButtonController{},"*:LinkButton")
	beego.Router("/menubutton",&controllers.ButtonController{},"*:MenuButton")
}
