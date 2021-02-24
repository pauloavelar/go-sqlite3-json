package utils

import "reflect"

func IsValueInArray(array []interface{}, value interface{}) bool {
	for _, each := range array {
		if reflect.DeepEqual(each, value) {
			return true
		}
	}
	return false
}
