package mathexp

import "reflect"

func isTypeAllowed(array []reflect.Type, val reflect.Type) bool {
	if array == nil || len(array) == 0 {
		return false
	}
	for _, typ := range array {
		if typ == val {
			return true
		}
	}
	return false
}
