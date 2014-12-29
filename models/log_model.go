package models

import (
	"time"
)

type Log struct {
	Id       int `orm:"pk"`
	UserName string
	Content  string
	CreateOn time.Time `orm:"auto_now_add;type(datetime)"`
}
