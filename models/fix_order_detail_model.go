package models

import (
	"time"
)

type FixOrderDetail struct {
	Id       int `orm:"pk"`
	OrderId  int
	Detail   string
	Images   string
	CreateOn time.Time `orm:"auto_now_add;type(datetime)"`
	CreateBy string
}
