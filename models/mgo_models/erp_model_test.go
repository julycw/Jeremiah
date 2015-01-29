package mgo_models

import (
	"gopkg.in/mgo.v2/bson"
	"testing"
)

type Book struct {
	ID_   bson.ObjectId `bson:"_id"`
	Title string        `bson:"title"`
	Price float32       `bson:"price"`
}

func TestGetERPContext(t *testing.T) {
	ctx, _ := GetERPContext("BookStore", Book{})
	if ctx == nil {
		t.Error("error!")
	}

	book := Book{
		ID_:   bson.NewObjectId(),
		Title: "C PLUS PLUS Primer",
		Price: 12.3,
	}

	//insert
	if err := ctx.Insert(&book); err != nil {
		t.Error(err.Error())
	}

	//find all
	var books []Book = make([]Book, 0)
	if err := ctx.Find(nil, &books); err != nil {
		t.Error(err.Error())
		return
	}

	//find one
	func() {
		var oneBook Book

		objectId := books[len(books)-1].ID_
		//by object id
		if err := ctx.FindId(objectId, &oneBook); err != nil {
			t.Error(err.Error())
		}

		stringId := books[len(books)-1].ID_.Hex()
		//by string id
		if err := ctx.FindId(stringId, &oneBook); err != nil {
			t.Error(err.Error())
		}

		if oneBook.Title == "" {
			t.Error("Result field is empty")
		}
	}()

	//mult-conditions searching

}
