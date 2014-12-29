package models

import (
	"time"
)

type FollowUser struct {
	Id            int `orm:"pk"`
	UserId        string
	WeixinId      string
	Name          string
	UserName      string
	NickName      string
	Phone         string
	Sex           string
	Description   string
	FollowOn      time.Time
	UnFollowOn    time.Time
	LastMessageOn time.Time
}
