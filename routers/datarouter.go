package routers

import (
	"github.com/astaxie/beego"
	"quickstart/controllers"
)

func init() {
	beego.Router("/sampledata",&controllers.DataController{},"*:SampleData")
	beego.Router("/getdata",&controllers.DataController{},"*:DataGrid")
	beego.Router("/selectdata",&controllers.DataController{},"*:SelectData")
	beego.Router("/searchdata",&controllers.DataController{},"*:SearchData")
	beego.Router("/complextool",&controllers.DataController{},"*:ComplexTool")
	beego.Router("/freezedata",&controllers.DataController{},"*:FreezeData")
	beego.Router("/changedata",&controllers.DataController{},"*:ChangeData")
	beego.Router("/formatdata",&controllers.DataController{},"*:FormatData")
	beego.Router("/sortdata",&controllers.DataController{},"*:SortData")
	beego.Router("/checkboxdata",&controllers.DataController{},"*:CheckBoxData")
	beego.Router("/paginationdata",&controllers.DataController{},"*:PaginationData")
	beego.Router("/editinlinedata",&controllers.DataController{},"*:EditInlineData")
	beego.Router("/calculatedata",&controllers.DataController{},"*:CalculateData")
	beego.Router("/mergedata",&controllers.DataController{},"*:MergeData")
	beego.Router("/imagedata",&controllers.DataController{},"*:ImageData")
	//数据底部
	beego.Router("/footdata",&controllers.DataController{},"*:FootData")
	beego.Router("/footgrid",&controllers.DataController{},"*:FootGrid")
	//
	beego.Router("/backgrounddata",&controllers.DataController{},"*:BackgroundData")
	beego.Router("/virtualscrolldata",&controllers.DataController{},"*:VirtualScrollData")

}
