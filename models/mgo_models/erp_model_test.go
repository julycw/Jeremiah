package mgo_models

import (
	"testing"
)

type Book struct {
	Title string  `bson:"title"`
	Price float32 `bson:"price"`
}

func TestGetERPContext(t *testing.T) {
	ctx, _ := GetERPContext("BookStore", Book{})
	if ctx == nil {
		t.Error("error!")
	}

	err := ctx.Insert(&Book{
		"C PLUS PLUS Primer",
		12.3,
	})

	if err != nil {
		t.Error(err.Error())
	}
}
