package web

import (
	"github.com/julycw/Jeremiah/controllers/manage"
	"github.com/julycw/Jeremiah/models"
)

type NewsController struct {
	manage.BaseController
}

func (this *NewsController) Get() {
	this.CheckAvaliable("文章管理")
	this.TplNames = "manage/web/news.tpl"
	this.LayoutSections["LayoutHeader"] = "manage/web/news_header.tpl"
	this.LayoutSections["LayoutFooter"] = "manage/web/news_script.tpl"

	var newsList []models.News
	if _, err := models.Orm.QueryTable("News").All(&newsList); err == nil {
		this.Data["newsList"] = &newsList
	}
}

func (this *NewsController) Post() {
	this.CheckAvaliable("文章管理")
	this.Redirect("/manage/web/news", 302)
}

func (this *NewsController) Add() {
	this.CheckAvaliable("文章管理")
	category, _ := this.GetInt("category")
	title := this.GetString("title")
	html := this.GetString("html")
	state, _ := this.GetInt("state")
	tags := this.GetString("tags")

	new_news := models.News{
		Category: int(category),
		Title:    title,
		Html:     html,
		Tags:     tags,
		State:    int(state),
		CreateBy: this.User.UserName,
		Click:    0,
	}

	if _, err := models.Orm.Insert(&new_news); err == nil {

	}
	this.Redirect("/manage/web/news", 302)
}

func (this *NewsController) Modify() {
	this.CheckAvaliable("文章管理")
	id, _ := this.GetInt("id")
	category, _ := this.GetInt("category")
	title := this.GetString("title")
	state, _ := this.GetInt("state")
	html := this.GetString("html")
	tags := this.GetString("tags")

	modify_news := models.News{
		Id:       int(id),
		Category: int(category),
		Title:    title,
		Html:     html,
		State:    int(state),
		Tags:     tags,
	}

	if _, err := models.Orm.Update(&modify_news); err == nil {

	}
	this.Redirect("/manage/web/news", 302)
}

// @router /api/manage/web/news/getNewsById [post]
func (this *NewsController) GetNewsById() {
	this.CheckAvaliable("文章管理")
	id, _ := this.GetInt("id")

	var news models.News
	news.Id = int(id)
	if err := models.Orm.Read(&news); err == nil {
		this.Data["json"] = &news
	}

	this.ServeJson()
}

// @router /api/manage/web/news/deleteNewsById [post]
func (this *NewsController) DeleteNewsById() {
	this.CheckAvaliable("文章管理")
	id, _ := this.GetInt("id")

	if _, err := models.Orm.Delete(&models.News{Id: int(id)}); err == nil {
		this.Data["json"] = "success"
	} else {
		this.Data["json"] = "err:" + err.Error()
	}

	this.ServeJson()
}
