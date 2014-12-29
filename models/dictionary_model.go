package models

import (
	"reflect"
)

type Dictionary struct {
	Id    int `orm:"pk"`
	Item  string
	Key   string
	Value string
	Name  string
}

type Dict int

func (this Dict) IntValue() int64 {
	value := reflect.ValueOf(this)
	return value.Int()
}

//新闻状态
type NewsState Dict

const (
	EnumNewsStateReady NewsState = iota + 1
	EnumNewsStatePublished
	EnumNewsStateHidden
)

func (this NewsState) String() string {
	switch this {
	case EnumNewsStateReady:
		return "准备发布"
	case EnumNewsStatePublished:
		return "已发布"
	case EnumNewsStateHidden:
		return "不可见"
	}
	return "未知"
}

//新闻类型
type NewsCategory Dict

const (
	EnumNewsCategoryWhatWeDo NewsCategory = iota + 1
	EnumNewsCategoryHot
	EnumNewsCategoryNormal
)

func (this NewsCategory) String() string {
	switch this {
	case EnumNewsCategoryWhatWeDo:
		return "我们做什么"
	case EnumNewsCategoryHot:
		return "最新热点"
	case EnumNewsCategoryNormal:
		return "一般新闻"
	}
	return "未知"
}

//产品类型
type ProductType Dict

const (
	EnumProductTypeNormal ProductType = iota + 1
)

func (this ProductType) String() string {
	switch this {
	case EnumProductTypeNormal:
		return "智能家居"
	}
	return "未知"
}
