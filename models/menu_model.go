package models

type MenuModel struct {
	Id        int `orm:"pk"`
	ParentId  int
	MenuTitle string
	Action    string
}
