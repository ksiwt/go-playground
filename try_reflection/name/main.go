package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

// reflect.Name returns the type's name within its package for a defined type.
// For other (non-defined) types it returns the empty string.
func main() {
	rt := reflect.TypeOf(100)
	fmt.Println(rt.Name()) //=> int

	rt = reflect.TypeOf([]int{})
	fmt.Println(rt.Name()) //=>

	type T1 struct {
		F1 int
		F2 string
	}

	rt = reflect.TypeOf(T1{})
	fmt.Println(rt.Name()) //=> T1

	rt = reflect.TypeOf(struct{
		F1 int
		F2 string
	}{})
	fmt.Println(rt.Name()) //=>

	rt = reflect.TypeOf(rand.New(rand.NewSource(99)))
	fmt.Println(rt.Name()) //=>
}