package weixin

import (
	"github.com/astaxie/beego"
	"github.com/julycw/Jeremiah/cache"
	"github.com/julycw/Jeremiah/models"
	"log"
	"strings"
)

type JsonRet struct {
	Result  bool
	Message string
}

type DayFixOrderList struct {
	Day          string
	FixOrderList []models.FixOrderEx
}

type WeixinPageController struct {
	beego.Controller
}

func (this *WeixinPageController) Prepare() {
	this.Layout = "weixin/layout.tpl"
}

// @router /weixin/user [get]
func (this *WeixinPageController) UserInfo() {
	this.TplNames = "weixin/userInfo.tpl"
	if userID := this.GetString("userID"); userID != "" {
		var user models.FollowUser
		user.UserId = userID
		if err := models.Orm.Read(&user, "UserId"); err != nil {
			log.Println(err.Error())
		} else {
			this.Data["User"] = &user
		}
	}
}

// @router /weixin/user/save [post]
func (this *WeixinPageController) UserInfoSave() {
	if this.IsAjax() {
		id, _ := this.GetInt("id")
		name := this.GetString("name")
		phone := this.GetString("phone")
		sex := this.GetString("sex")

		user := models.FollowUser{
			Id:    int(id),
			Name:  name,
			Phone: phone,
			Sex:   sex,
		}

		var jsonRet JsonRet

		if num, err := models.Orm.Update(&user, "Name", "Phone", "Sex"); err != nil {
			log.Println(err.Error())
			jsonRet.Result = false
			jsonRet.Message = err.Error()
		} else {
			if num > 0 {
				jsonRet.Result = true
			} else {
				jsonRet.Result = false
				jsonRet.Message = "未修改任何信息"
			}
		}

		this.Data["json"] = &jsonRet
		this.ServeJson()

	} else {
		this.Redirect("http://www.google.com/", 302)
		return
	}
}

// @router /weixin/viewOrder [get]
func (this *WeixinPageController) ViewOrder() {
	this.TplNames = "weixin/viewOrder.tpl"
	var userID string
	if this.GetString("userID") != "" {
		userID = this.GetString("userID")
		this.Data["userID"] = this.GetString("userID")
	}
	var orderList []models.FixOrder

	if _, err := models.Orm.QueryTable("fix_order").Filter("UserId", userID).OrderBy("-create_on").Limit(5).All(&orderList); err != nil {
		log.Println(err.Error())
	} else {
		var dayFixOrderList []DayFixOrderList = make([]DayFixOrderList, 0)
		for _, order := range orderList {
			var orderEx models.FixOrderEx
			orderEx.Id = order.Id
			orderEx.OrderNumber = order.OrderNumber
			orderEx.OrderShortNumber = order.OrderShortNumber
			orderEx.OrderStatus = order.OrderStatus
			orderEx.UserId = order.UserId
			orderEx.Description = order.Description
			orderEx.CreateOn = order.CreateOn
			orderEx.CreateBy = order.CreateBy

			var orderDetails []models.FixOrderDetail
			models.Orm.QueryTable("fix_order_detail").Filter("order_id", order.Id).OrderBy("-create_on").All(&orderDetails)

			orderEx.OrderDetails = orderDetails

			dayStr := order.CreateOn.Format("01月02日")
			hasThisDay := false
			for i, day := range dayFixOrderList {
				if dayStr == day.Day {
					dayFixOrderList[i].FixOrderList = append(dayFixOrderList[i].FixOrderList, orderEx)
					hasThisDay = true
				}
			}

			if !hasThisDay {
				var dayFixOrders DayFixOrderList
				dayFixOrders.Day = dayStr
				dayFixOrders.FixOrderList = make([]models.FixOrderEx, 0)
				dayFixOrders.FixOrderList = append(dayFixOrders.FixOrderList, orderEx)
				dayFixOrderList = append(dayFixOrderList, dayFixOrders)
			}
		}
		this.Data["dayFixOrderList"] = &dayFixOrderList
	}
}

// @router /weixin/bindOrderNumber [get,post]
func (this *WeixinPageController) BindOrderNumber() {
	var userID string
	if this.GetString("userID") != "" {
		userID = this.GetString("userID")
		this.Data["userID"] = this.GetString("userID")
	}

	method := strings.ToLower(this.Ctx.Request.Method)
	switch method {
	case "get":
		this.TplNames = "weixin/bindOrderNumber.tpl"
	case "post":
		this.TplNames = "weixin/bindOrderNumberResult.tpl"
		if order_suffix := this.GetString("order_suffix"); len(order_suffix) == 4 {
			// cache.GlobalCache.Put(order_suffix, "20141114134433", 60)
			cacheKey := "ON:" + order_suffix
			if cache.GlobalCache.IsExist(cacheKey) {
				order_prefix := cache.GlobalCache.Get(cacheKey)
				cache.GlobalCache.Delete(cacheKey)

				var orderPair models.FixOrderPair
				orderPair.UserId = userID
				if err := models.Orm.Read(&orderPair, "UserId"); err != nil {
					orderPair.OrderNumber = order_prefix.(string) + "-" + order_suffix
					if _, err := models.Orm.Insert(&orderPair); err != nil {
						log.Println(err.Error())
						return
					}
				} else {
					orderPair.OrderNumber = order_prefix.(string) + "-" + order_suffix
					if _, err := models.Orm.Update(&orderPair, "OrderNumber"); err != nil {
						log.Println(err.Error())
						return
					}
				}
				this.Data["success"] = "success"
				this.Data["orderNumber"] = orderPair.OrderNumber
				this.Data["orderShortNumber"] = order_prefix.(string)[6:10] + order_suffix
				return
			}
		}
	}
}

// @router /weixin/error [get]
func (this *WeixinPageController) ShowError() {
	this.TplNames = "weixin/error.tpl"
}
