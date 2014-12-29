package models

import (
	"time"
)

type Note struct {
	Id           int `orm:"pk"`
	Content      string
	UserName     string
	Tags         string
	ParentNoteId int
	CreateOn     time.Time `orm:"auto_now_add;type(datetime)"`
}
