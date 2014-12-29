package models

type Option struct {
	Id          int `orm:"pk"`
	Section     string
	SectionName string
	Key         string
	Title       string
	Value       string
	Description string
}

type OptionSection struct {
	SectionName string
	Options     []Option
	HasNext     bool
}
