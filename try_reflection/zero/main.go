package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	rv := reflect.Zero(reflect.TypeOf(1))
	fmt.Println(rv) //=> 0

	type S1 struct {
		F1 string
		F2 bool
	}
	rv = reflect.Zero(reflect.TypeOf(S1{}))
	fmt.Println(rv) //=> { false}

	rv = reflect.Zero(reflect.TypeOf(time.Now()))
	fmt.Println(rv) //=> 0001-01-01 00:00:00 +0000 UTC
}