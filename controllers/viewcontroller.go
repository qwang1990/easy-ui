package controllers

import (
	"github.com/astaxie/beego"
	"time"
	"github.com/astaxie/beego/httplib"
	"fmt"
)

type User struct {
	Id    int         `form:"-"`
	Name  interface{} `form:"username"`
	Age   int         `form:"age,text,年龄："`
	Sex   string
	Intro string `form:",textarea"`
}


type ViewController struct {
	beego.Controller
}

func (this *ViewController) Get() {

	req := httplib.Get("http://beego.me/")
	str, _ := req.String()
	fmt.Println(str)


	pages := []struct {
		Num int
	}{{10}, {20}, {30}}

	item := []struct {
		Name int
		Age int
	}{{100,10}, {200,20}, {300,30}}

	this.Data["Total"] = 100
	this.Data["Pages"] = pages
	this.Data["Items"] = item

	this.Data["Maaaap"] = map[string]string{"name": "Beego"}


	ss :=[]string{"a","b","c"}
	this.Data["s"]=ss

	var mp = map[string]string{}
	mp["name"]="astaxie"
	mp["nickname"] = "haha"
	this.Data["m"]=mp

	this.Data["Form"] = &User{}

	this.Data["Time"] = time.Now()
	this.TplName="viewtest.tpl"
}
