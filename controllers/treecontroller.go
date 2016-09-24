package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type TreeController struct {
	beego.Controller
}

func (this *TreeController) SampleTree() {
	this.TplName = "tree/sampletree.html"
}

func (this *TreeController) AsyncTree() {
	this.TplName = "tree/asynctree.html"
}

type Node struct {
	Id int `json:"id"`
	Text string `json:"text"`
	State string `json:"state"`
}

type Result struct {
	Id int `json:"id"`
	Nodes []*Node `json:"children"`
	Text string `json:"text"`
	State string `json:"state"`
} 

func (this *TreeController) TreeData() {
	nodes := make([] *Node,3)
	for i := 0; i<len(nodes); i++ {
		nodes[i] = &Node{Id:i+1,Text:"wang"}
	}

	for i := 0; i<len(nodes); i++ {
		fmt.Println(nodes[i])
	}

	//result := &Result{Nodes:nodes,Id:0,Text:"heheh",State:"open"}

	result := make([] *Result,1)

	result[0] = &Result{Nodes:nodes,Id:0,Text:"heheh",State:"open"}

	this.Data["json"] = &result
	this.ServeJSON()

}

func (this *TreeController) AddNodeTree() {
	this.TplName = "tree/addnodetree.html"
}


func (this *TreeController) DragTree() {
	this.TplName = "tree/dragtree.html"
}



type Diy struct {
	Id int `json:"id"`
	ParentId int `json:"parentId"`
	Name string `json:"name"`
}

func (this *TreeController) DiyData() {
	diys := make([]*Diy,10)
	diys[0] = &Diy{Id:1,ParentId:0,Name:"Foods"}
	diys[1] = &Diy{Id:2,ParentId:1,Name:"Fruits"}
	diys[2] = &Diy{Id:3,ParentId:1,Name:"Vegetables"}
	diys[3] = &Diy{Id:4,ParentId:2,Name:"apple"}
	diys[4] = &Diy{Id:5,ParentId:2,Name:"orange"}
	diys[5] = &Diy{Id:6,ParentId:3,Name:"tomato"}
	diys[6] = &Diy{Id:7,ParentId:3,Name:"carrot"}
	diys[7] = &Diy{Id:8,ParentId:3,Name:"cabbage"}
	diys[8] = &Diy{Id:9,ParentId:3,Name:"potato"}
	diys[9] = &Diy{Id:10,ParentId:3,Name:"lettuce"}

	this.Data["json"] = &diys
	this.ServeJSON()
}

func (this *TreeController) DiyLoadTree() {
	this.TplName = "tree/diyloadtree.html"
}






