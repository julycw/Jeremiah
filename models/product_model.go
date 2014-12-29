package models

import (
	"time"
)

type Product struct {
	Id          int `orm:"pk"`
	Title       string
	Type        int
	Tags        string
	Description string
	Html        string
	Images      string
	Image       string
	Display     string
	Sort        int
	Hot         string
	Visible     string
	CreateOn    time.Time `orm:"auto_now_add;type(datetime)"`
	CreateBy    string
	Click       int
}
