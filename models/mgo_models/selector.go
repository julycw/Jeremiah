package mgo_models

import (
	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// "reflect"
	"strings"
)

// type Condition struct {
// 	fieldName string
// 	operator  string
// 	params    []interface{}
// 	SelectorMaker
// }

type SelectorMaker struct {
	selector *bson.M
}

func (this *SelectorMaker) And(params ...interface{}) *SelectorMaker {
	lenParams := len(params)
	for lenParams > 0 {
		//handle condition
		// if reflect.TypeOf(params[0]).Kind() == reflect.Struct {
		// 	for _, c := range params {
		// 		cond := c.(Condition)
		// 		(*this.selector)[cond.fieldName] = cond.Selector()
		// 	}
		// 	break
		// }

		//handle others
		var fieldName, operation string
		fieldName = params[0].(string)
		// if length of params is 2,means get params without operation sign, just set it operation as $eq
		if lenParams == 2 {
			operation = "$eq"
			(*this.selector)[fieldName] = params[1]
		} else if lenParams >= 3 {
			operation = params[1].(string)
			if !strings.HasPrefix(operation, "$") {
				operation = "$" + operation
			}
			switch operation {
			case "$eq":
				(*this.selector)[fieldName] = params[1]
			case "$ne", "$gt", "$lt", "$gte", "$lte":
				if m, has := (*this.selector)[fieldName]; has {
					m.(bson.M)[operation] = params[2]
					(*this.selector)[fieldName] = m
				} else {
					(*this.selector)[fieldName] = bson.M{operation: params[2]}
				}
			case "$in", "$nin":
				(*this.selector)[fieldName] = bson.M{operation: params[2:]}
			case "$regex":
				/**
				 * Valid options as of this writing are 'i' for case insensitive matching,
				 * 'm' for multi-line matching, 'x' for verbose mode, 'l' to make \w, \W,
				 * and similar be locale-dependent, 's' for dot-all mode (a '.' matches everything),
				 * and 'u' to make \w, \W, and similar match unicode.
				 */
				(*this.selector)[fieldName] = bson.M{operation: bson.RegEx{params[2].(string), "iu"}}
			case "$elemMatch":
				(*this.selector)[fieldName] = bson.M{operation: bson.M{"$eq": params[2]}}
			default:
				(*this.selector)[fieldName] = bson.M{operation: params[2]}
			}
		}

		break
	}
	return this
}

func (this *SelectorMaker) Selector() *bson.M {
	return this.selector
}

func NewSelector() *SelectorMaker {
	return &SelectorMaker{
		selector: &bson.M{},
	}
}
