package web

import (
	// "encoding/json"
	"fmt"
	"github.com/julycw/Jeremiah/controllers/manage"
	"github.com/julycw/Jeremiah/models/mgo_models"
	"github.com/julycw/Jeremiah/module/parser"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type ERPController struct {
	manage.BaseController
}

func (this *ERPController) Index() {
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

	var computer mgo_models.Computer
	if url != "" {
		this.Data["url"] = url
		if ret, err := parser.Parse(url); err != nil {
			this.Data["Error"] = fmt.Sprintf("%s", err.Error())
			computer.ID = bson.NewObjectId()
			computer.IDStr = computer.ID.Hex()
		} else {
			computer = ret.(mgo_models.Computer)
			if this.checkExsit(computer.Code) {
				this.Data["Warning"] = fmt.Sprintf("编号为%s的商品已存在，如果提交，将对数据进行覆盖", computer.Code)
			}
		}
	}

	this.Data["Form"] = &computer
}

// @router /api/manage/web/erp [get]
func (this *ERPController) Get() {
	selector := mgo_models.NewSelector()
	model := this.GetString("model")
	if model != "" && model != "all" {
		selector.And("model", "regex", model)
	}
	priceMin, _ := this.GetInt("suggestPriceMin")
	if priceMin >= 0 {
		selector.And("suggestPrice", "gte", priceMin)
	}
	priceMax, _ := this.GetInt("suggestPriceMax")
	if priceMax >= 0 {
		selector.And("suggestPrice", "lte", priceMax)
	}
	cpu := this.GetString("cpu")
	if cpu != "" && cpu != "all" {
		selector.And("cpuModel", "regex", cpu)
	}
	graphic := this.GetString("graphic")
	if graphic != "" && graphic != "all" {
		selector.And("graphicsMemorySize", "regex", graphic)
	}
	memory := this.GetString("memory")
	if memory != "" && memory != "all" {
		selector.And("memorySize", "regex", memory)
	}
	disk := this.GetString("disk")
	if disk != "" && disk != "all" {
		selector.And("diskSize", "regex", disk)
	}
	os := this.GetString("os")
	if os != "" && os != "all" {
		selector.And("os", "regex", os)
	}

	if ctx, err := mgo_models.GetERPContext("jeremiah", "computer", mgo_models.Computer{}); err != nil {
		log.Println(err.Error())
	} else {
		count := ctx.Count(selector.Selector())
		results := make([]mgo_models.Computer, count)
		if findErr := ctx.Find(selector.Selector(), &results); findErr != nil {
			log.Println(findErr.Error())
		} else {
			this.Data["json"] = &results
		}
	}

	this.ServeJson()
}

// @router /api/manage/web/erp [post]
func (this *ERPController) Add() {
	computer := mgo_models.Computer{}
	if err := this.ParseForm(&computer); err != nil {
		this.ResponseError(err)
	}

	if ctx, err := mgo_models.GetERPContext("jeremiah", "computer", mgo_models.Computer{}); err != nil {
		log.Println(err.Error())
		this.ResponseError(err)
	} else {

		if this.checkExsit(computer.Code) { //update
			//无需更改ID
			computer.ID = bson.ObjectId("")
			computer.IDStr = ""

			updateErr := ctx.Update(mgo_models.NewSelector().And("code", computer.Code).Selector(), &computer)
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

func (this *ERPController) checkExsit(code string) bool {
	exsits := false
	if ctx, err := mgo_models.GetERPContext("jeremiah", "computer", mgo_models.Computer{}); err != nil {
		log.Println(err.Error())
	} else {
		selector := mgo_models.NewSelector().And("code", code)
		results := make([]mgo_models.Computer, 0)
		if findError := ctx.Find(selector.Selector(), &results); findError != nil {
			log.Println(findError.Error())
		} else {
			if len(results) >= 1 {
				exsits = true
			}
		}
	}

	return exsits
}
