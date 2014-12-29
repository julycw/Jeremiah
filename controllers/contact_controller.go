package controllers

type ContactController struct {
	BaseController
}

func (this *ContactController) Get() {
	this.TplNames = "web/contact.tpl"
}
