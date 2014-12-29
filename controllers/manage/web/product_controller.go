package web

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/julycw/Jeremiah/controllers/manage"
	"github.com/julycw/Jeremiah/models"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type ProductController struct {
	manage.BaseController
}

func (this *ProductController) Get() {
	this.CheckAvaliable("商品管理")
	this.TplNames = "manage/web/product.tpl"
	this.LayoutSections["LayoutHeader"] = "manage/web/product_header.tpl"
	this.LayoutSections["LayoutFooter"] = "manage/web/product_script.tpl"

	var productList []models.Product
	if _, err := models.Orm.QueryTable("Product").All(&productList); err == nil {
		this.Data["productList"] = &productList
	}
}

func (this *ProductController) Post() {
	this.CheckAvaliable("商品管理")
	this.Redirect("/manage/web/product", 302)
}

func (this *ProductController) Add() {
	this.CheckAvaliable("商品管理")
	title := this.GetString("title")
	_type, _ := this.GetInt("type")
	html := this.GetString("html")
	tags := this.GetString("tags")
	description := this.GetString("description")
	images := this.GetString("images")
	display := this.GetString("display")
	sort, _ := this.GetInt("sort")
	visible := this.GetString("visible")
	hot := this.GetString("hot")
	image := ""
	if picture, fileheader, err := this.GetFile("image"); err == nil {
		defer picture.Close()
		if strings.HasSuffix(fileheader.Filename, ".jpg") {
			//保存文件
			dir, err := os.Stat("./static/img/product/")
			if err != nil || !dir.IsDir() {
				os.Mkdir("./static/img/product/", os.ModeDir)
			}

			// 保存的文件路径和名称
			var filemd5name string
			if bytes, err := ioutil.ReadAll(picture); err == nil {
				md5Maker := md5.New()
				md5Maker.Write(bytes)
				filemd5name = hex.EncodeToString(md5Maker.Sum(nil))
				filedir := "./static/img/product/" + filemd5name + ".jpg"
				if newfile, err := os.Create(filedir); err == nil {
					defer newfile.Close()
					if _, err := newfile.Write(bytes); err == nil {
						image = filedir[1:]
					} else {
						log.Println(err.Error())
					}
				} else {
					log.Println(err.Error())
				}
			} else {
				log.Println(err.Error())
			}
		}
	}

	new_product := models.Product{
		Title:       title,
		Type:        int(_type),
		Html:        html,
		Tags:        tags,
		Description: description,
		Images:      images,
		Image:       image,
		Display:     display,
		Sort:        int(sort),
		Hot:         hot,
		Visible:     visible,
		CreateBy:    this.User.UserName,
		Click:       0,
	}

	if _, err := models.Orm.Insert(&new_product); err == nil {

	}
	this.Redirect("/manage/web/product", 302)
}

func (this *ProductController) Modify() {
	this.CheckAvaliable("商品管理")
	id, _ := this.GetInt("id")
	title := this.GetString("title")
	_type, _ := this.GetInt("type")
	html := this.GetString("html")
	tags := this.GetString("tags")
	description := this.GetString("description")
	images := this.GetString("images")
	image := this.GetString("image")
	display := this.GetString("display")
	sort, _ := this.GetInt("sort")
	hot := this.GetString("hot")
	visible := this.GetString("visible")

	modify_product := models.Product{
		Id:          int(id),
		Title:       title,
		Type:        int(_type),
		Html:        html,
		Tags:        tags,
		Description: description,
		Images:      images,
		Image:       image,
		Display:     display,
		Sort:        int(sort),
		Hot:         hot,
		Visible:     visible,
	}

	if _, err := models.Orm.Update(&modify_product); err == nil {

	}
	this.Redirect("/manage/web/product", 302)
}

// @router /api/manage/web/product/getProductById [post]
func (this *ProductController) GetProductById() {
	this.CheckAvaliable("商品管理")
	id, _ := this.GetInt("id")

	var product models.Product
	product.Id = int(id)
	if err := models.Orm.Read(&product); err == nil {
		this.Data["json"] = &product
	}

	this.ServeJson()
}

// @router /api/manage/web/product/deleteProductById [post]
func (this *ProductController) DeleteProductById() {
	this.CheckAvaliable("商品管理")
	id, _ := this.GetInt("id")

	if _, err := models.Orm.Delete(&models.Product{Id: int(id)}); err == nil {
		this.Data["json"] = "success"
	} else {
		this.Data["json"] = "err:" + err.Error()
	}

	this.ServeJson()
}
