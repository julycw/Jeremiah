package models

import (
	"time"
)

type User struct {
	Id         int `orm:"pk"`
	UserName   string
	PassWord   string
	NickName   string
	RoleId     int
	RegisterOn time.Time `orm:"auto_now_add;type(datetime)"`
	LoginOn    time.Time
	LogoutOn   time.Time
}
