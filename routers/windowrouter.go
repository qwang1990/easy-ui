package routers

import (
	"github.com/astaxie/beego"
	"quickstart/controllers"
)

func init() {
	beego.Router("/win",&controllers.WinController{},"*:Win")
	beego.Router("/diywin",&controllers.WinController{},"*:DiyWin")
	beego.Router("/layoutwin",&controllers.WinController{},"*:LayoutWin")
	beego.Router("/dialogwin",&controllers.WinController{},"*:DialogWin")


}
