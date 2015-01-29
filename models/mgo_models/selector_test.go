package mgo_models

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestSelectorMaker(t *testing.T) {
	selectorMaker := NewSelector()
	selectorMaker.And("title", "hello")
	selectorMaker.And("price", "$gt", 11)
	selectorMaker.And([]interface{}{"type", "$in", "小说", "散文", "诗歌"}...)
	selectorMaker.And("content", "$regex", "hello world")

	selector := selectorMaker.Selector()

	selector2 := bson.M{
		"title": "hello",
		"price": bson.M{
			"$gt": 11,
		},
		"type": bson.M{
			"$in": []string{"小说", "散文", "诗歌"},
		},
		"content": bson.M{
			"$regex": bson.RegEx{"hello world", "u"},
		},
	}

	bytes, _ := json.Marshal(selector)
	bytes2, _ := json.Marshal(&selector2)
	hasError := false
	if len(bytes) != len(bytes2) {
		t.Error("Result not equal!")
		hasError = true
	} else {
		for i := 0; i < len(bytes); i++ {
			if bytes[i] != bytes2[i] {
				t.Error("Result not equal!")
				hasError = true
				break
			}
		}
	}

	if hasError {
		fmt.Println("by maker:", string(bytes))
		fmt.Println("by mgo.M:", string(bytes2))
	}

}
