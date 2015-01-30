package web

import (
	"github.com/julycw/Jeremiah/controllers/manage"
	// "github.com/julycw/Jeremiah/models/mgo_models"
	"github.com/julycw/Jeremiah/module"
	"strings"
)

type ERPController struct {
	manage.BaseController
}

func (this *ERPController) Prepare() {
	this.Layout = "manage/layout.tpl"
}

func (this *ERPController) Get() {
	// this.CheckAvaliable("文章管理")
	this.TplNames = "manage/web/erp.tpl"
	this.LayoutSections["LayoutHeader"] = "manage/web/erp_header.tpl"
	this.LayoutSections["LayoutFooter"] = "manage/web/erp_script.tpl"
}

func (this *ERPController) Scan() {
	this.TplNames = "manage/web/erp_scan.tpl"

	url := this.GetString("url")
	if strings.Contains(url, "jd.com") {
		ret, err := module.ParseJD(url)
		if err != nil {
			this.Data["Error"] = err.Error()
		} else {
			this.Data["Content"] = ret.(string)
		}
	}
}

// @router /api/manage/web/erp/add [post]
func (this *ERPController) Add() {

}
