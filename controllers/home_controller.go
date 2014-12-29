package controllers

import (
	"github.com/julycw/Jeremiah/models"
	"log"
)

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {
	this.TplNames = "web/home.tpl"
	this.Data["Title"] = "慈溪耶利米 - 智能家居、家庭影院"
	this.LayoutSections["LayoutHead"] = "web/home_slider.tpl"
	this.LayoutSections["LayoutFooter"] = "web/home_script.tpl"

	text := models.Option{}

	if err := models.Orm.QueryTable("option").Filter("Key", "custom_made").One(&text); err == nil {
		this.Data["custom_made"] = text.Value
	} else {
		log.Println(err.Error())
	}
	if err := models.Orm.QueryTable("option").Filter("Key", "service").One(&text); err == nil {
		this.Data["service"] = text.Value
	} else {
		log.Println(err.Error())
	}
	if err := models.Orm.QueryTable("option").Filter("Key", "success_fse").One(&text); err == nil {
		this.Data["success_case"] = text.Value
	} else {
		log.Println(err.Error())
	}
	if err := models.Orm.QueryTable("option").Filter("Key", "welcome").One(&text); err == nil {
		this.Data["welcome"] = text.Value
	} else {
		log.Println(err.Error())
	}

	var whatWeDoList []models.News
	if _, err := models.Orm.QueryTable("news").Filter("category", models.EnumNewsCategoryWhatWeDo).Filter("state", models.EnumNewsStatePublished).Limit(10).All(&whatWeDoList); err == nil {
		this.Data["whatWeDoList"] = &whatWeDoList
	}

	var hotProductList []models.Product
	if _, err := models.Orm.QueryTable("product").Filter("hot", 1).Filter("visible", 1).OrderBy("-Sort").Limit(6).All(&hotProductList); err == nil {
		this.Data["hotProductList"] = &hotProductList
	}

	var productList []models.Product
	if _, err := models.Orm.QueryTable("product").Filter("hot", 0).Filter("visible", 1).OrderBy("-Sort").Limit(6).All(&productList); err == nil {
		this.Data["productList"] = &productList
	}

}
