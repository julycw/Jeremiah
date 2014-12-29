package controllers

import (
	"github.com/julycw/Jeremiah/models"
)

type GalleryController struct {
	BaseController
}

func (this *GalleryController) Get() {
	this.TplNames = "web/gallery.tpl"
	var productList []models.Product
	if _, err := models.Orm.QueryTable("product").Filter("visible", 1).OrderBy("-Sort").Limit(24).All(&productList); err == nil {
		this.Data["productList"] = &productList
	}
}

func (this *GalleryController) Detail() {
	this.TplNames = "web/product.tpl"
	id, _ := this.GetInt(":id")

	var product models.Product
	if err := models.Orm.QueryTable("product").Filter("id", id).Filter("visible", 1).One(&product); err == nil {
		this.Data["product"] = &product
	} else {
		this.Abort("404")
		this.StopRun()
	}
}
