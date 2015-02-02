package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	this.Layout = "web/layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.Data["SiteAddr"] = beego.AppConfig.String("webaddr")
	this.Data["SitePort"] = beego.AppConfig.String("httpport")
}

func (this *BaseController) Get() {
	this.Abort("404")
}

func (this *BaseController) Post() {
	this.Abort("404")
}

func init() {

}
