package utils

import "reflect"

func IsNil(value interface{}) bool {
	return value == nil || (IsNillable(value) && reflect.ValueOf(value).IsNil())
}

func IsNillable(value interface{}) bool {
	kind := reflect.ValueOf(value).Kind()
	return kind == reflect.Ptr || kind == reflect.Slice
}
