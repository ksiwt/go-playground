package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := 1
	fmt.Println(reflect.ValueOf(v).CanSet()) //=> false

	fmt.Println(reflect.ValueOf(&v).Elem().CanSet()) //=> true

	type S1 struct {
		F1 int
		f2 bool
	}
	vv := S1{F1: 100, f2: true}
	fmt.Println(reflect.ValueOf(&vv).Elem().CanSet()) //=> true
	fmt.Println(reflect.ValueOf(&vv).Elem().FieldByName("F1").CanSet()) //=> true
	fmt.Println(reflect.ValueOf(&vv).Elem().FieldByName("f2").CanSet()) //=> false
}
