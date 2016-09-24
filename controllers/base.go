package controllers

import "github.com/astaxie/beego"

type RController struct {
beego.Controller
}

type Temp struct {
	Name string
	Age int
}

func (this *RController) Prepare() {
	temp := &Temp{"haahah",2}
	this.Data["xml"] = temp

	this.ServeXML()
	this.StopRun()
}