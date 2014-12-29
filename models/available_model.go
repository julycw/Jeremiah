package models

type Avaliable struct {
	Id             int `orm:"pk"`
	Avaliable      string
	ControllerName string
	ActionName     string
}
