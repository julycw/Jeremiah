package models

type Role struct {
	Id         int `orm:"pk"`
	RoleId     int
	RoleName   string
	Avaliables string
}
