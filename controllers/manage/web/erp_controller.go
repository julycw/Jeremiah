package web

import (
	"github.com/julycw/Jeremiah/controllers/manage"
	"github.com/julycw/Jeremiah/models/mgo_models"
	"github.com/julycw/Jeremiah/module"
	"gopkg.in/mgo.v2/bson"
	"log"
	"strings"
)

type ERPController struct {
	manage.BaseController
}

func (this *ERPController) Get() {
	// this.CheckAvaliable("文章管理")
	this.TplNames = "manage/web/erp.tpl"
	this.LayoutSections["LayoutHeader"] = "manage/web/erp_header.tpl"
	this.LayoutSections["LayoutFooter"] = "manage/web/erp_script.tpl"
}

func (this *ERPController) Scan() {
	this.TplNames = "manage/web/erp_scan.tpl"
	this.LayoutSections["LayoutHeader"] = "manage/web/erp_scan_header.tpl"
	this.LayoutSections["LayoutFooter"] = "manage/web/erp_scan_script.tpl"

	url := this.GetString("url")
	this.Data["url"] = url

	var computer mgo_models.Computer
	if strings.Contains(url, "jd.com") {
		if ret, err := module.ParseJD(url); err != nil {
			this.Data["Error"] = err.Error()
			computer.ID = bson.NewObjectId()
			computer.IDStr = computer.ID.Hex()
		} else {
			computer = ret.(mgo_models.Computer)
		}
	}
	this.Data["Form"] = &computer
}

// @router /api/manage/web/erp/add [post]
func (this *ERPController) Add() {
	computer := mgo_models.Computer{}
	if err := this.ParseForm(&computer); err != nil {
		this.ResponseError(err)
	}

	//put _id fiedls
	computer.ID = bson.ObjectIdHex(computer.IDStr)

	if ctx, err := mgo_models.GetERPContext("jeremiah", "computer", mgo_models.Computer{}); err != nil {
		log.Println(err.Error())
		this.ResponseError(err)
	} else {
		insertErr := ctx.Insert(&computer)
		if insertErr != nil {
			log.Println(insertErr.Error())
			this.ResponseError(insertErr)
		} else {
			this.Data["json"] = "success"
		}
	}

	this.ServeJson()
}
