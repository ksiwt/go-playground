package main

import (
	"fmt"
	"reflect"
)

// IsZero reports whether v is the zero value for its type.
// It panics if the argument is invalid.
func main() {
	var v string
	fmt.Println(reflect.ValueOf(v).IsZero()) //=> true
	v = ""
	fmt.Println(reflect.ValueOf(v).IsZero()) //=> true
	v = "a"
	fmt.Println(reflect.ValueOf(v).IsZero()) //=> false
}
