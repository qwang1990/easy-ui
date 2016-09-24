package routers

import (
	"github.com/astaxie/beego"
	"quickstart/controllers"
)

func init() {
	beego.Router("/sampletree",&controllers.TreeController{},"*:SampleTree")
	beego.Router("/asynctree",&controllers.TreeController{},"*:AsyncTree")

	beego.Router("/treedata",&controllers.TreeController{},"*:TreeData")
	beego.Router("/addnodetree",&controllers.TreeController{},"*:AddNodeTree")
	beego.Router("/dragtree",&controllers.TreeController{},"*:DragTree")
	beego.Router("/diyloadtree",&controllers.TreeController{},"*:DiyLoadTree")

	beego.Router("/diydata",&controllers.TreeController{},"*:DiyData")
}
