package models

import (
	"time"
)

type FixOrder struct {
	Id               int `orm:"pk"`
	OrderNumber      string
	OrderShortNumber string
	OrderStatus      int
	UserId           string
	Description      string
	CreateOn         time.Time `orm:"auto_now_add;type(datetime)"`
	CreateBy         string
}

type FixOrderEx struct {
	FixOrder
	OrderDetails []FixOrderDetail
}

type OrderStatus int

const (
	OrderStatusCreated OrderStatus = iota + 1
	OrderStatusProcessing
	OrderStatusClosed
	OrderStatusFinished
)

func (this OrderStatus) String() string {
	switch this {
	case OrderStatusCreated:
		return "已创建"
	case OrderStatusProcessing:
		return "进行中"
	case OrderStatusClosed:
		return "已废弃"
	case OrderStatusFinished:
		return "已结束"
	default:
		return "未知"
	}
}
