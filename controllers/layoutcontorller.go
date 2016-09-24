package controllers

import "github.com/astaxie/beego"

type LayoutController struct {
	beego.Controller
}

func (this *LayoutController) SampleLayout()  {
	this.TplName = "layout/samplelayout.html"
}

func (this *LayoutController) MNS()  {
	this.TplName = "layout/mns.html"
}

func (this *LayoutController) Fold()  {
	this.TplName = "layout/fold.html"
}

func (this *LayoutController) Tab()  {
	this.TplName = "layout/tab.html"
}


func (this *LayoutController) AutoAdd()  {
	this.TplName = "layout/autoadd.html"
}

func (this *LayoutController) Xp()  {
	this.TplName = "layout/xp.html"
}