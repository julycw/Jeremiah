package mgo_models

import (
	"errors"
	"labix.org/v2/mgo"
	"reflect"
)

var sess *mgo.Session = nil

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

func (this *ERPContext) checkData(data interface{}) bool {
	t := reflect.TypeOf(data)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t != this.modelType {
		return false
	}

	return true
}
func (this *ERPContext) checkDatas(data []interface{}) bool {
	for _, v := range data {
		if !this.checkData(v) {
			return false
		}
	}

	return true
}

func (this *ERPContext) Insert(data ...interface{}) error {

	if !this.checkDatas(data) {
		return errors.New("Error data struct type")
	}

	return this.db.C(this.collectionName).Insert(data...)
}

func GetERPContext(collection string, modelType interface{}) (*ERPContext, error) {
	t := reflect.TypeOf(modelType)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return nil, errors.New("Unknown model type")
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
