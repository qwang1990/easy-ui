package routers

import (
	"github.com/astaxie/beego"
	"quickstart/controllers"
)

func init() {
	beego.Router("/layoutsample",&controllers.LayoutController{},"*:SampleLayout")
	beego.Router("/mns",&controllers.LayoutController{},"*:MNS")
	beego.Router("/fold",&controllers.LayoutController{},"*:Fold")
	beego.Router("/tab",&controllers.LayoutController{},"*:Tab")
	beego.Router("/autoadd",&controllers.LayoutController{},"*:AutoAdd")
	beego.Router("/xp",&controllers.LayoutController{},"*:Xp")
}
