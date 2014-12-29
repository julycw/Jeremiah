package models

type FixOrderPair struct {
	Id          int `orm:"pk"`
	UserId      string
	OrderNumber string
}
