package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"net/http"
	"os"
	"io"
)

type FormController struct {
	beego.Controller
}

func (this *FormController) AjaxForm() {
	this.TplName="form/ajaxform.html"
}

func (this *FormController) FormAjaxData() {
	fmt.Println(this.Ctx.Request)
	this.Data["json"]="hehe"
	this.ServeJSON()
}

func (this *FormController) ValidateForm() {
	this.TplName="form/validateform.html"
}

func (this *FormController) ComboTreeForm() {
	this.TplName="form/combotreeform.html"
}

func (this *FormController) FormatComboForm() {
	this.TplName="form/formatcomboform.html"
}

type FtData struct {
	Id int `json:"id"`
	Text string `json:"text"`
	Icon string `json:"icon"`

}
func (this *FormController) FormatData() {
	ftdatas := make([]*FtData,7)
	ftdatas[0] = &FtData{Id:1,Text:"Word",Icon:"doc.png"}
	ftdatas[1] = &FtData{Id:2,Text:"Excel",Icon:"xls.png"}
	ftdatas[2] = &FtData{Id:3,Text:"Zip",Icon:"zip.png"}
	ftdatas[3] = &FtData{Id:4,Text:"Html",Icon:"html.png"}
	ftdatas[4] = &FtData{Id:5,Text:"Css",Icon:"css.png"}
	ftdatas[5] = &FtData{Id:6,Text:"Text",Icon:"txt.png"}
	ftdatas[6] = &FtData{Id:7,Text:"PowerPoint",Icon:"ppt.png"}


	for i := 0; i<len(ftdatas); i++ {
		fmt.Println(ftdatas[i])
	}

	this.Data["json"] = &ftdatas
	this.ServeJSON()


}


func (this *FormController) ComboGridForm() {
	this.TplName="form/combogridform.html"
}



type UploadResponse struct {
	Error int `json:"error"`
	Url string `json:"url"`
}

func (this *FormController) UploadFile() {
	fmt.Println(this.Ctx.Request)
	files, fileheader,err:= this.GetFile("imgFile")
	if err != nil {
		http.Error(this.Ctx.ResponseWriter, err.Error(), http.StatusNoContent)
		return
	}
	dst, err := os.Create("upload/" + fileheader.Filename)
	defer dst.Close()
	if err != nil {
		http.Error(this.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	//copy the uploaded file to the destination file
	if _, err := io.Copy(dst, files); err != nil {
		http.Error(this.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	uprs := &UploadResponse{Error:0,Url:"/files/"+fileheader.Filename}

	this.Data["json"] = &uprs
	this.ServeJSON()
}









