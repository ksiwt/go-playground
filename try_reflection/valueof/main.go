package main

import (
	"fmt"
	"reflect"
)

func plus(v1, v2 interface{}) interface{} {
	// ValueOf returns a new Value initialized to the concrete value
	// stored in the interface i. ValueOf(nil) returns the zero Value.
	rv1, rv2 := reflect.ValueOf(v1), reflect.ValueOf(v2)
	if isInt(rv1.Kind()) && isInt(rv2.Kind()) {
		return interface{}(rv1.Int() + rv2.Int())
	}
	if isString(rv1.Kind()) && isString(rv2.Kind()) {
		return interface{}(rv1.String() + rv2.String())
	}
	return nil
}

func isInt(k reflect.Kind) bool {
	return k == reflect.Int || k == reflect.Int8 ||
		k == reflect.Int16 || k == reflect.Int32 ||
		k == reflect.Int64
}

func isString(k reflect.Kind) bool {
	return k == reflect.String
}

func main() {
	fmt.Println(plus(100, 200)) //=> 300
	fmt.Println(plus("100", "200")) //=> 10200
	fmt.Println(plus(100, "200")) //=> nil
}