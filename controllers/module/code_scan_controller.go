package module

import (
	"github.com/astaxie/beego"
)

type CodeScanController struct {
	beego.Controller
}

// @router /api/manage/weixin/fix/scanFixOrder [get]
func (this *CodeScanController) ScanFixOrder() {
	this.TplNames = "manage/weixin/order-scaned.tpl"
	orderNumber := this.GetString("order")
	this.Data["OrderNumber"] = orderNumber
}
