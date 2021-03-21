package main

import (
	"fmt"
	"reflect"
)

// Elem returns the value that the interface v contains
// or that the pointer v points to.
// It panics if v's Kind is not Interface or Ptr.
// It returns the zero Value if v is nil.
func main() {
	i := 1
	// 元の値を更新するためには当該のアドレスを知っている必要がある
	// Elemはinterfaceの値もしくはポインタの参照先を返す
	// なので以下のようにValueOfにポインタを渡して、Elemでポインタが指している先の値を取得する
	rv := reflect.ValueOf(&i).Elem()
	rv.SetInt(2)
	fmt.Println(i) //=> 2

	b := false
	rv = reflect.ValueOf(&b).Elem()
	rv.SetBool(true)
	fmt.Println(b) //=> true

	s := []int{10, 11, 12}
	rv = reflect.ValueOf(&s).Elem()
	rv.Set(reflect.ValueOf([]int{20, 21, 22}))
	fmt.Println(s) //=> [20, 21, 22]
}
