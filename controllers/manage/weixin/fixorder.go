package weixin

import (
	"fmt"
	"github.com/julycw/Jeremiah/cache"
	"github.com/julycw/Jeremiah/controllers/manage"
	"github.com/julycw/Jeremiah/models"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type DayFixOrderList struct {
	Day          string
	FixOrderList []models.FixOrder
}

type FixOrderController struct {
	manage.BaseController
}

func makeOrderNumber() string {
	orderNumber_P1 := time.Now().Format("20060102150405")
	orderNumber_P2 := strconv.FormatInt(1000+rand.Int63n(9000), 10)
	cache.GlobalCache.Put("ON:"+orderNumber_P2, orderNumber_P1, 3600)
	return fmt.Sprintf("%s-%s", orderNumber_P1, orderNumber_P2)
}

func (this *FixOrderController) Get() {
	this.CheckAvaliable("订单管理")
	this.TplNames = "manage/weixin/order.tpl"
	this.LayoutSections["LayoutFooter"] = "manage/weixin/order_script.tpl"

	var orderList []models.FixOrder

	this.Data["orderNumber"] = makeOrderNumber()

	if _, err := models.Orm.QueryTable("fix_order").OrderBy("-create_on").Limit(100).All(&orderList); err != nil {
		log.Println(err.Error())
	} else {
		var dayFixOrderList []DayFixOrderList = make([]DayFixOrderList, 0)
		for _, note := range orderList {
			dayStr := note.CreateOn.Format("01月02日")
			hasThisDay := false
			for i, day := range dayFixOrderList {
				if dayStr == day.Day {
					dayFixOrderList[i].FixOrderList = append(dayFixOrderList[i].FixOrderList, note)
					hasThisDay = true
				}
			}

			if !hasThisDay {
				var dayFixOrders DayFixOrderList
				dayFixOrders.Day = dayStr
				dayFixOrders.FixOrderList = make([]models.FixOrder, 0)
				dayFixOrders.FixOrderList = append(dayFixOrders.FixOrderList, note)
				dayFixOrderList = append(dayFixOrderList, dayFixOrders)
			}
		}
		this.Data["dayFixOrderList"] = &dayFixOrderList
	}
}

func (this *FixOrderController) Post() {

}

// @router /api/manage/weixin/fix/getOrderNumber [post]
func (this *FixOrderController) GetOrderNumber() {
	this.CheckAvaliable("订单管理")
	this.ResponseData(makeOrderNumber())
}

// @router /api/manage/weixin/fix/getFixOrderDetailsByOrderId [post]
func (this *FixOrderController) GetFixOrderDetailsByOrderId() {
	this.CheckAvaliable("订单管理")
	orderId, _ := this.GetInt("id")
	var orderDetails []models.FixOrderDetail
	if _, err := models.Orm.QueryTable("fix_order_detail").Filter("order_id", orderId).OrderBy("-create_on").All(&orderDetails); err != nil {
		log.Println(err.Error())
	}
	this.Data["json"] = &orderDetails

	this.ServeJson()
}

// @router /api/manage/weixin/fix/addFixOrder [post]
func (this *FixOrderController) AddFixOrder() {
	this.CheckAvaliable("订单管理")
	userID := this.GetString("userID")
	description := this.GetString("description")
	orderNumber := this.GetString("orderNumber")
	var fixOrder models.FixOrder
	fixOrder.OrderNumber = orderNumber
	fixOrder.OrderShortNumber = orderNumber[6:10] + orderNumber[15:]
	fixOrder.OrderStatus = int(models.OrderStatusCreated)
	fixOrder.UserId = userID
	fixOrder.Description = description
	fixOrder.CreateBy = this.User.UserName

	if id, err := models.Orm.Insert(&fixOrder); err != nil {
		log.Println(err.Error())
	} else {
		fixOrder.Id = int(id)
	}

	this.Data["json"] = &fixOrder
	this.ServeJson()
}

// @router /api/manage/weixin/fix/addFixOrderDetail [post]
func (this *FixOrderController) AddFixOrderDetail() {
	this.CheckAvaliable("订单管理")
	orderID, _ := this.GetInt("id")
	detail := this.GetString("detail")
	var orderDetail models.FixOrderDetail
	orderDetail.OrderId = int(orderID)
	orderDetail.Detail = detail
	orderDetail.CreateOn = time.Now()
	orderDetail.CreateBy = this.User.UserName

	if detail != "" {
		if num, err := models.Orm.Insert(&orderDetail); err != nil {
			log.Println(err.Error())
		} else {
			orderDetail.Id = int(num)
		}
	}

	this.Data["json"] = &orderDetail
	this.ServeJson()
}

// @router /api/manage/weixin/fix/getOrderBindedUser [post]
func (this *FixOrderController) GetOrderBindedUser() {
	this.CheckAvaliable("订单管理")
	var followUser models.FollowUser
	if orderNumber := this.GetString("orderNumber"); orderNumber != "" {
		var pair models.FixOrderPair
		pair.OrderNumber = orderNumber
		if err := models.Orm.Read(&pair, "OrderNumber"); err != nil {
			log.Println(err.Error())
		} else {
			followUser.UserId = pair.UserId
			if err := models.Orm.Read(&followUser, "UserId"); err != nil {
				log.Println(err.Error())
			} else {
				models.Orm.Delete(&pair)
			}
		}
	}
	this.Data["json"] = &followUser

	this.ServeJson()
}

func init() {
	rand.Seed(time.Now().Unix())
}
