package controllers

import "github.com/astaxie/beego"

type ButtonController struct {
	beego.Controller
}

func (this *ButtonController) SampleMenu()  {
	this.TplName = "buttonandmenu/samplemenu.html"
}


func (this *ButtonController) LinkButton()  {
	this.TplName = "buttonandmenu/linkbutton.html"
}

func (this *ButtonController) MenuButton()  {
	this.TplName = "buttonandmenu/menubutton.html"
}