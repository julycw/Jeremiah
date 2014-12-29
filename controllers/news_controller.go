package controllers

import (
	"github.com/julycw/Jeremiah/models"
)

type NewsController struct {
	BaseController
}

func (this *NewsController) Get() {
	this.TplNames = "web/newsList.tpl"
	var newsList []models.News
	if _, err := models.Orm.QueryTable("news").Filter("state", models.EnumNewsStatePublished).Limit(50).All(&newsList); err == nil {
		this.Data["newsList"] = &newsList
	}
}

func (this *NewsController) Detail() {
	this.TplNames = "web/news.tpl"
	id, _ := this.GetInt(":id")

	var news models.News
	if err := models.Orm.QueryTable("news").Filter("id", id).Filter("state", models.EnumNewsStatePublished).One(&news); err == nil {
		this.Data["news"] = &news
	} else {
		this.Abort("404")
		this.StopRun()
	}
}

func (this *NewsController) WhatWeDo() {
	this.TplNames = "web/whatWeDo.tpl"
}
