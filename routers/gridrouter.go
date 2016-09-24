package routers

import (
	"github.com/astaxie/beego"
	"quickstart/controllers"
)

func init() {
	beego.Router("/datagrid",&controllers.EasyController{},"*:DataGrid")
	beego.Router("/data",&controllers.EasyController{})
	beego.Router("/editgrid",&controllers.EasyController{},"*:EditDatagrid")
	beego.Router("/save",&controllers.EasyController{},"*:Delete")
	beego.Router("/detailgrid",&controllers.EasyController{},"*:DetailDatagrid")
	beego.Router("/rss",&controllers.EasyController{},"*:Rss")
}
