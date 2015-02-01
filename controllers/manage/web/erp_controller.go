package web

import (
	"fmt"
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

	if ctx, err := mgo_models.GetERPContext("jeremiah", "computer", mgo_models.Computer{}); err != nil {
		log.Println(err.Error())
		this.ResponseError(err)
	} else {
		this.Data["Count"] = ctx.Count(nil)
	}

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
			this.Data["Error"] = fmt.Sprintf("%s%s", err.Error(), " 请手动录入")
			computer.ID = bson.NewObjectId()
			computer.IDStr = computer.ID.Hex()
		} else {
			computer = ret.(mgo_models.Computer)
			if this.checkExsitJDNumber(computer.JDNumber) {
				this.Data["Warning"] = fmt.Sprintf("京东编号为%s的商品已存在，如果提交，将对数据进行覆盖", computer.JDNumber)
			}
		}
	} else {
		this.Data["Error"] = "暂不支持扫描非京东数据，请手动录入"
	}
	this.Data["Form"] = &computer
}

// @router /api/manage/web/erp/add [post]
func (this *ERPController) Add() {
	computer := mgo_models.Computer{}
	if err := this.ParseForm(&computer); err != nil {
		this.ResponseError(err)
	}

	if ctx, err := mgo_models.GetERPContext("jeremiah", "computer", mgo_models.Computer{}); err != nil {
		log.Println(err.Error())
		this.ResponseError(err)
	} else {

		if this.checkExsitJDNumber(computer.JDNumber) { //update
			//无需更改ID
			computer.ID = bson.ObjectId("")
			computer.IDStr = ""

			updateErr := ctx.Update(mgo_models.NewSelector().And("jdNumber", computer.JDNumber).Selector(), &computer)
			if updateErr != nil {
				log.Println(updateErr.Error())
				this.ResponseError(updateErr)
			} else {
				this.Data["json"] = "success"
			}

		} else { //insert
			//设置ID
			computer.ID = bson.ObjectIdHex(computer.IDStr)

			insertErr := ctx.Insert(&computer)
			if insertErr != nil {
				log.Println(insertErr.Error())
				this.ResponseError(insertErr)
			} else {
				this.Data["json"] = "success"
			}
		}

	}

	this.ServeJson()
}

func (this *ERPController) checkExsitJDNumber(jdNumber string) bool {
	exsits := false
	if ctx, err := mgo_models.GetERPContext("jeremiah", "computer", mgo_models.Computer{}); err != nil {
		fmt.Println(err.Error())
	} else {
		selector := mgo_models.NewSelector().And("jdNumber", jdNumber)
		results := make([]mgo_models.Computer, 0)
		if findError := ctx.Find(selector.Selector(), &results); findError != nil {
			fmt.Println(findError.Error())
		} else {
			if len(results) >= 1 {
				fmt.Println(len(results))
				exsits = true
			}
		}
	}

	return exsits
}
