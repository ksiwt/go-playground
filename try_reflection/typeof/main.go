package main

import (
	"fmt"
	"reflect"
)

func isInt(v interface{}) bool {
	// TypeOf returns the reflection Type that represents the dynamic type of i.
	// If i is a nil interface value, TypeOf returns nil.
	rt := reflect.TypeOf(v)
	switch rt.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	default:
		return false
	}
}

func main() {
	fmt.Println(isInt(1)) //=> true
	fmt.Println(isInt("1")) //=> false
	fmt.Println(isInt(false)) //=> false
}