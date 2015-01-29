package mgo_models

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"reflect"
)

var sess *mgo.Session = nil

var (
	ErrorWrongDataType    error = errors.New("Wrong data struct type")
	ErrorModelIsNotStruct error = errors.New("Model type must be struct")
)

const (
	ip     = "localhost"
	port   = "27017"
	dbName = "jeremiah"
)

type ERPContext struct {
	modelType      reflect.Type
	collectionName string
	db             *mgo.Database
}

func (this *ERPContext) giveMeConfidence(data interface{}) bool {
	t := reflect.TypeOf(data)
	//Pointer、Slice、Array、Channel等类型调用Elem后可获得基本数据类型
	for t.Kind() == reflect.Ptr || t.Kind() == reflect.Slice {
		t = t.Elem()
	}

	if t != this.modelType {
		return false
	}

	return true
}
func (this *ERPContext) giveMeConfidences(data []interface{}) bool {
	for _, v := range data {
		if !this.giveMeConfidence(v) {
			return false
		}
	}

	return true
}

func (this *ERPContext) FindId(id interface{}, result interface{}) error {
	if !this.giveMeConfidence(result) {
		return ErrorWrongDataType
	}

	var objectId bson.ObjectId
	t := reflect.TypeOf(id)
	if t == reflect.TypeOf(objectId) {
		objectId = id.(bson.ObjectId)
	} else if t.Kind() == reflect.String {
		objectId = bson.ObjectIdHex(id.(string))
	}

	return this.db.C(this.collectionName).Find(bson.M{"_id": objectId}).One(result)
}

func (this *ERPContext) Find(selector, results interface{}) error {
	t := reflect.TypeOf(results)
	if t.Kind() != reflect.Ptr {
		return ErrorWrongDataType
	}
	t = t.Elem()
	if t.Kind() != reflect.Slice {
		return ErrorWrongDataType
	}
	if !this.giveMeConfidence(results) {
		return ErrorWrongDataType
	}

	return this.db.C(this.collectionName).Find(selector).All(results)
}

func (this *ERPContext) Update(selector, update interface{}) error {
	if !this.giveMeConfidence(update) {
		return ErrorWrongDataType
	}

	return this.db.C(this.collectionName).Update(selector, update)
}

func (this *ERPContext) Insert(data ...interface{}) error {
	if !this.giveMeConfidences(data) {
		return ErrorWrongDataType
	}

	return this.db.C(this.collectionName).Insert(data...)
}

func GetERPContext(collection string, modelType interface{}) (*ERPContext, error) {
	t := reflect.TypeOf(modelType)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return nil, ErrorModelIsNotStruct
	}
	return &ERPContext{
		collectionName: collection,
		db:             sess.DB(dbName),
		modelType:      t,
	}, nil
}

func init() {
	var err error
	if sess, err = mgo.Dial(ip + ":" + port); err != nil {
		panic("Initial mongodb connect failed..." + ip + ":" + port)
	}

	sess.SetMode(mgo.Monotonic, true)
}
