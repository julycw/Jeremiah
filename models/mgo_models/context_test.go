package mgo_models

import (
	"gopkg.in/mgo.v2/bson"
	"testing"
)

type Book struct {
	ID_    bson.ObjectId `bson:"_id"`
	Number int           `bson:"number"`
	Title  string        `bson:"title"`
	Price  float32       `bson:"price"`
	Tag    []string      `bson:"tag"`
	Type   string        `bson:"type"`
	Author Author        `bson:"author"`
}

type Author struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

var books []Book = []Book{
	Book{
		Title: "阿里扎的山",
		Price: 25,
		Type:  "散文",
		Tag:   []string{"Hot", "adult"},
	},
	Book{
		Title: "黑客的修养",
		Price: 172.3,
		Type:  "装逼",
		Tag:   []string{"Hot"},
	},
	Book{
		Title: "乔丹",
		Price: 23.1,
		Type:  "传记",
		Tag:   []string{"NBA"},
	},
	Book{
		Title:  "詹姆斯",
		Price:  23.2,
		Type:   "传记",
		Tag:    []string{"NBA"},
		Author: Author{"Wangwang", 19},
	},
	Book{
		Title: "库里",
		Price: 30,
		Type:  "儿歌",
		Tag:   []string{"NBA", "children"},
	},
	Book{
		Title: "杜兰特",
		Price: 35,
		Type:  "传记",
		Tag:   []string{"NBA"},
	},
}

func TestGetERPContext(t *testing.T) {
	DialDB("127.0.0.1", "27017")
	ctx, _ := GetERPContext("jeremiah_test", "BookStore", Book{})
	if ctx == nil {
		t.Error("error!")
	}

	//insert something make sure db exist.
	books[0].ID_ = bson.NewObjectId()
	books[0].Number = 0

	//insert
	if err := ctx.Insert(&books[0]); err != nil {
		t.Error(err.Error())
	}
	//drop
	if err := ctx.Drop(); err != nil {
		t.Error(err.Error())
	}

	for i := range books {
		books[i].ID_ = bson.NewObjectId()
		books[i].Number = i + 1
		//insert
		if err := ctx.Insert(&books[i]); err != nil {
			t.Error(err.Error())
		}
	}

	//test Count()
	if count := ctx.Count(nil); count != len(books) {
		t.Error("Error result of Count()")
	}

	//find all
	var booksAll []Book = make([]Book, 0)
	if err := ctx.Find(nil, &booksAll); err != nil {
		t.Error(err.Error())
		return
	}

	//find one
	func() {
		var oneBook Book

		objectId := booksAll[len(booksAll)-1].ID_
		//by object id
		if err := ctx.FindId(objectId, &oneBook); err != nil {
			t.Error(err.Error())
		}

		stringId := booksAll[len(booksAll)-1].ID_.Hex()
		//by string id
		if err := ctx.FindId(stringId, &oneBook); err != nil {
			t.Error(err.Error())
		}

		if oneBook.Title == "" {
			t.Error("Result field is empty")
		}
	}()

	//mult-conditions searching
	selector := NewSelector()
	selector.And("tag", "$elemMatch", "NBA").And("price", "$gte", 30.0)

	if count := ctx.Count(selector.Selector()); count != 2 {
		t.Errorf("Count should be 2,but return %d", count)
	}

	bookResult := make([]Book, 0)
	if err := ctx.Find(selector.Selector(), &bookResult); err != nil {
		t.Error(err.Error())
	}

	//update
	selector = NewSelector()
	selector.And("type", "$ne", "传记")
	if err := ctx.Update(selector.Selector(), &Book{
		Price:  99.0,
		Tag:    []string{"haha", "hehe"},
		Author: Author{"July", 18},
	}); err != nil {
		t.Error(err.Error())
	}

	func() {
		var resultBooks []Book = make([]Book, 0)
		if err := ctx.Limit(1).Find(selector.Selector(), &resultBooks); err != nil {
			t.Error(err.Error())
		} else {
			if resultBooks[0].Price != 99 {
				t.Errorf("Price should be 99,but return %f", resultBooks[0].Price)
			}
		}

	}()

	//delete
	if del, err := ctx.Delete(selector.Selector()); err != nil {
		t.Error(err.Error())
	} else {
		if del != 3 {
			t.Errorf("error:%d", del)
		}
	}

}
