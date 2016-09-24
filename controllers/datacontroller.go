package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"time"
	"strconv"
)

type Item struct {
	Itemid string
	Productid string
	Listprice float64
	Unitcost float64
	Attr1 string
	Status string
	Date time.Time
	Amount float64
}

type Goods struct {
	Id int
	Name string
	Price float64
}

type FooterData struct {
	Gs []*Goods	`json:"rows"`
	Fs []*Footer	`json:"footer"`
}

type Footer struct {
	Name string``
	Price float64
}

type DataController struct {
	beego.Controller
}

func (this *DataController) SampleData() {
	this.TplName = "data/sampledata.html"
}


func (this *DataController) DataGrid() {

	fmt.Println(this.Ctx.Request)

	items := make([] *Item,10)
	for i := 0; i<len(items); i++ {
		items[i] = &Item{Itemid:"EST-"+strconv.Itoa(i),Productid:"FI-SW-01",Listprice:100,Unitcost:1.5,Attr1:"hehe",Status:"P",Date:time.Now(),Amount:10}
	}

	for i := 0; i<len(items); i++ {
		fmt.Println(items[i])
	}

	this.Data["json"]=&items
	this.ServeJSON()

	//this.TplName="datagrid/datagrid.html"

}


func (this *DataController) FootGrid() {

	fmt.Println(this.Ctx.Request)

	goods := make([] *Goods,10)
	for i := 0; i<len(goods); i++ {
		goods[i] = &Goods{Id:i,Name:"hehe",Price:10.1}
	}

	for i := 0; i<len(goods); i++ {
		fmt.Println(goods[i])
	}

	footer := make([] *Footer,2)

	footer[0] = &Footer{Name:"Total",Price:101}
	footer[1] = &Footer{Name:"Average",Price:10.1}

	footerdata := &FooterData{Gs:goods,Fs:footer}



	this.Data["json"]=&footerdata
	this.ServeJSON()

	//this.TplName="datagrid/datagrid.html"

}

func (this *DataController) SelectData() {
	this.TplName = "data/selectdata.html"
}


func (this *DataController) SearchData() {
	this.TplName = "data/searchdata.html"
}

func (this *DataController) ComplexTool() {
	this.TplName = "data/complextool.html"
}

func (this *DataController) FreezeData() {
	this.TplName = "data/freezedata.html"
}

func (this *DataController) ChangeData() {
	this.TplName = "data/changedata.html"
}

func (this *DataController) FormatData() {
	this.TplName = "data/formatdata.html"
}

func (this *DataController) SortData() {
	this.TplName = "data/sortdata.html"
}


func (this *DataController) CheckBoxData() {
	this.TplName = "data/checkboxdata.html"
}

func (this *DataController) PaginationData() {
	this.TplName = "data/paginationdata.html"
}

func (this *DataController) EditInlineData() {
	this.TplName = "data/editinlinedata.html"
}

func (this *DataController) CalculateData() {
	this.TplName = "data/calculatedata.html"
}

func (this *DataController) MergeData() {
	this.TplName = "data/mergedata.html"
}

func (this *DataController) ImageData() {
	this.TplName = "data/imagedata.html"
}

func (this *DataController) FootData() {
	this.TplName = "data/footdata.html"
}

func (this *DataController) BackgroundData() {
	this.TplName = "data/backgrounddata.html"
}

func (this *DataController) VirtualScrollData() {
	this.TplName = "data/virtualscrolldata.html"
}

//type Info struct {
//	Inv int `json:"inv"`
//	Date time.Time `json:"date"`
//	Name string `json:"name"`
//}













