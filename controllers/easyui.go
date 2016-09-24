package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type EasyController struct {
	beego.Controller
}

//func (c *MainController) Get() {
//	c.Data["Website"] = "beego.me"
//	c.Data["Email"] = "astaxie@gmail.com"
//	c.TplName = "index.tpl"
//}

type Wang struct {
	Firstname string
	Lastname string
	Phone string
	Email string
}

func (this *EasyController) DataGrid() {
	users := make([]*Wang,10)
	for i := 0; i<len(users); i++ {
		users[i] = &Wang{Firstname:"qi",Lastname:"wang",Phone:"1388888888",Email:"hhh"}
	}

	for i := 0; i<len(users); i++ {
		fmt.Println(users[i])
	}

	this.Data["json"]=&users
	this.ServeJSON()

	//this.TplName="datagrid/datagrid.html"

}

func (this *EasyController) Get() {
	this.TplName="datagrid/datagrid.html"
}

func (this *EasyController) EditDatagrid() {
	this.TplName="datagrid/editdatagrid.html"
}


func (this *EasyController) Delete() {
	firstname := this.GetString("Firstname");
	fmt.Println(firstname+"hehehehe")
	fmt.Println(this.Ctx.Request)
	//this.Data["json"]="success"
	this.TplName="datagrid/editdatagrid.html"
}

func (this *EasyController) DetailDatagrid() {
	this.TplName="datagrid/detailgrid.html"
}

func (this *EasyController) Rss() {
	this.TplName="datagrid/rssfeed.html"
}







