package controllers

type ServicesController struct {
	BaseController
}

func (this *ServicesController) Get() {
	this.TplNames = "web/services.tpl"
}
