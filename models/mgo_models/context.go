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

var (
	defaultLimit = 30
	defaultSkip  = 0
)

type ERPContext struct {
	modelType      reflect.Type
	collectionName string
	limit          int
	skip           int
	sort           []string
	db             *mgo.Database
}

func (this *ERPContext) giveMeConfidence(data interface{}) bool {
	t := reflect.TypeOf(data)
	//Pointer、Slice、Array、Channel等类型调用Elem后可获得基本数据类型
	for t.Kind() == reflect.Ptr || t.Kind() == reflect.Slice || t.Kind() == reflect.Map {
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

func (this *ERPContext) Limit(params ...int) *ERPContext {
	paramsNum := len(params)
	if paramsNum >= 1 {
		this.limit = params[0]
		if paramsNum >= 2 {
			this.skip = params[1]
		}
	}

	return this
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
	defer this.Limit(defaultLimit, defaultSkip)
	defer this.Sort([]string{}...)

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

	query := this.db.C(this.collectionName).Find(selector).Skip(this.skip).Limit(this.limit)
	if len(this.sort) > 0 {
		query.Sort(this.sort...)
	}
	err := query.All(results)
	if err != nil {
		return err
	}
	return nil
}

func (this *ERPContext) Count(selector interface{}) int {
	count, _ := this.db.C(this.collectionName).Find(selector).Count()
	return count
}

func (this *ERPContext) Update(selector, update interface{}) error {
	if !this.giveMeConfidence(update) {
		return ErrorWrongDataType
	}

	m := bson.M{}

	//将update转换为bson.M类型
	value := reflect.ValueOf(update)
	for value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	for i := 0; i < value.NumField(); i++ {
		if bsonField := this.modelType.Field(i).Tag.Get("bson"); bsonField != "" {
			v := value.Field(i)
			for !v.CanInterface() {
				v = v.Elem()
			}

			//if a field of struct has default value, or the field is empty slice/array/map,just ignore it.
			if v.Kind() == reflect.Slice || v.Kind() == reflect.Array || v.Kind() == reflect.Map {
				if v.Cap() > 0 {
					m[bsonField] = v.Interface()
				}
			} else {
				if v.Interface() != reflect.Zero(v.Type()).Interface() {
					m[bsonField] = v.Interface()
				}
			}

		}
	}
	_, err := this.db.C(this.collectionName).UpdateAll(selector, bson.M{"$set": m})
	if err != nil {
		return err
	}
	return nil
}

func (this *ERPContext) Insert(data ...interface{}) error {
	if !this.giveMeConfidences(data) {
		return ErrorWrongDataType
	}

	return this.db.C(this.collectionName).Insert(data...)
}

func (this *ERPContext) Delete(selector interface{}) (int, error) {
	count, err := this.db.C(this.collectionName).Find(selector).Count()
	if err != nil {
		return 0, err
	}
	_, err = this.db.C(this.collectionName).RemoveAll(selector)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (this *ERPContext) Sort(fields ...string) {
	this.sort = fields
}

func (this *ERPContext) Drop() error {
	return this.db.C(this.collectionName).DropCollection()
}

func (this *ERPContext) Collection() *mgo.Collection {
	return this.db.C(this.collectionName)
}

func GetERPContext(database, collection string, dataModel interface{}) (*ERPContext, error) {
	if sess == nil {
		return nil, errors.New("Get context before register mongodb, please call DialDB() first.")
	}
	t := reflect.TypeOf(dataModel)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return nil, ErrorModelIsNotStruct
	}
	return &ERPContext{
		skip:           defaultSkip,
		limit:          defaultLimit,
		collectionName: collection,
		db:             sess.DB(database),
		modelType:      t,
	}, nil
}

func DialDB(ip, port string) {
	var err error
	if sess, err = mgo.Dial(ip + ":" + port); err != nil {
		panic("Initial mongodb connect failed..." + ip + ":" + port)
	}

	sess.SetMode(mgo.Monotonic, true)
}
