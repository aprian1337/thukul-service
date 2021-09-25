package helpers

import "reflect"

func IsZero(x interface{}) bool {
	return x == reflect.Zero(reflect.TypeOf(x)).Interface()
}
