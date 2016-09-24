package controllers

import "github.com/astaxie/beego"

type WinController struct {
	beego.Controller
}

func (this *WinController) Win() {
	this.TplName = "window/win.html"
}

func (this *WinController) DiyWin() {
	this.TplName = "window/diywin.html"
}

func (this *WinController) LayoutWin() {
	this.TplName = "window/layoutwin.html"
}

func (this *WinController) DialogWin() {
	this.TplName = "window/dialogwin.html"
}