package main

import (
	"fmt"
	"reflect"
)

// reflect.ValueOf returns a new Value initialized to the concrete value
// stored in the interface i. ValueOf(nil) returns the zero Value.
func main() {
	rv1 := reflect.ValueOf(100)
	rv2 := reflect.ValueOf(100)
	rv3 := reflect.ValueOf(300)

	fmt.Println(rv1.Interface() == rv2.Interface()) //=> true
	fmt.Println(rv1.Interface() == rv3.Interface()) //=> false

	rv4 := reflect.ValueOf("aaa")
	rv5 := reflect.ValueOf("aaa")
	rv6 := reflect.ValueOf("bbb")

	fmt.Println(rv4.Interface() == rv5.Interface()) //=> true
	fmt.Println(rv4.Interface() == rv6.Interface()) //=> false

	rv7 := reflect.ValueOf([]int{1, 2, 3})
	rv8 := reflect.ValueOf([]int{1, 2, 3})
	rv9 := reflect.ValueOf([]int{4, 5, 6})

	fmt.Println(rv7.Interface() == rv8.Interface()) //=> true
	fmt.Println(rv7.Interface() == rv9.Interface()) //=> false
}