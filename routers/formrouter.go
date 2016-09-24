package routers

import (
	"github.com/astaxie/beego"
	"quickstart/controllers"
)

func init() {
	beego.Router("/ajaxform",&controllers.FormController{},"*:AjaxForm")

	beego.Router("/formajaxdata",&controllers.FormController{},"*:FormAjaxData")
	beego.Router("/validateform",&controllers.FormController{},"*:ValidateForm")
	beego.Router("/combotreeform",&controllers.FormController{},"*:ComboTreeForm")
	beego.Router("/formatcomboform",&controllers.FormController{},"*:FormatComboForm")

	beego.Router("/combotdata",&controllers.FormController{},"*:FormatData")
	beego.Router("/combogridform",&controllers.FormController{},"*:ComboGridForm")
}
