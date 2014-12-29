package controllers

type AboutController struct {
	BaseController
}

func (this *AboutController) Get() {
	this.TplNames = "web/about.tpl"
}
