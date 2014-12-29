package models

import (
	"time"
)

type News struct {
	Id       int `orm:"pk"`
	Category int
	Title    string
	Html     string
	Tags     string
	State    int
	CreateOn time.Time `orm:"auto_now_add;type(datetime)"`
	CreateBy string
	Click    int
}
