package models

import (
	"time"
)

type Message struct {
	Id       int `orm:"pk"`
	UserId   string
	Content  string
	MsgType  string
	MsgId    string
	CreateOn time.Time `orm:"auto_now_add;type(datetime)"`
}

type MessageEx struct {
	Message
	User        FollowUser
	CreateOnFmt string
}
