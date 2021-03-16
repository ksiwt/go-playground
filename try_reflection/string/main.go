package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

// reflect.String returns a string representation of the type.
// The string representation may use shortened package names
// (e.g., base64 instead of "encoding/base64") and is not
// guaranteed to be unique among types. To test for type identity,
// compare the Types directly.
func main() {
	rt := reflect.TypeOf(100)
	fmt.Println(rt.String()) //=> int

	rt = reflect.TypeOf([]int{})
	fmt.Println(rt.String()) //=> []int

	type T1 struct {
		F1 int
		F2 string
	}

	rt = reflect.TypeOf(T1{})
	fmt.Println(rt.String()) //=> main.T1

	rt = reflect.TypeOf(struct{
		F1 int
		F2 string
	}{})
	fmt.Println(rt.String()) //=> struct { F1 int; F2 string }

	rt = reflect.TypeOf(rand.New(rand.NewSource(99)))
	fmt.Println(rt.String()) //=> *rand.Rand
}
