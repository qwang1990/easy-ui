package controllers

import (
	"github.com/astaxie/beego"
)

type TragController struct {
	beego.Controller
}


func (this *TragController) SampleDrag() {
	this.TplName="drag/sampledrag.html"
}



func (this *TragController) Shopping() {
	this.TplName="drag/shoppingcart.html"
}