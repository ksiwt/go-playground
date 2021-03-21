package main

import (
	"fmt"
	"reflect"
)

func main() {
	rv := reflect.ValueOf([]int{1, 3, 5, 7, 9})
	for i := 0; i < rv.Len(); i++ {
		el := rv.Index(i)
		fmt.Println(i, el.Interface())
	}
	//0 1
	//1 3
	//2 5
	//3 7
	//4 9
}
