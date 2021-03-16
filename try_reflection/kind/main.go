package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type S struct{}

// reflect.Kind returns the specific kind of this type.
func main() {
	fmt.Println(reflect.TypeOf(0).Kind() == reflect.Int)               //=> true
	fmt.Println(reflect.TypeOf(uint(0)).Kind() == reflect.Uint)        //=> true
	fmt.Println(reflect.TypeOf(float64(0)).Kind() == reflect.Float64)  //=> true
	fmt.Println(reflect.TypeOf(false).Kind() == reflect.Bool)          //=> true //
	fmt.Println(reflect.TypeOf("").Kind() == reflect.String)           //=> true
	fmt.Println(reflect.TypeOf([1]int{}).Kind() == reflect.Array)      //=> true
	fmt.Println(reflect.TypeOf([]int{}).Kind() == reflect.Slice)       //=> true
	fmt.Println(reflect.TypeOf(map[int]bool{}).Kind() == reflect.Map)  //=> true
	fmt.Println(reflect.TypeOf(make(chan int)).Kind() == reflect.Chan) //=> true
	fmt.Println(reflect.TypeOf(func() {}).Kind() == reflect.Func)      //=> true
	fmt.Println(reflect.TypeOf(S{}).Kind() == reflect.Struct)          //=> true
	fmt.Println(reflect.TypeOf(&S{}).Kind() == reflect.Ptr)            //=> true

	fmt.Println(AnyToFloat(100))
	fmt.Println(AnyToFloat(uint64(200)))
	fmt.Println(AnyToFloat(float64(300)))
	fmt.Println(AnyToFloat(complex128(400)))
}

func AnyToFloat(v interface{}) float64 {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64:
		return float64(rv.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64:
		return float64(rv.Int())
	case reflect.Float32, reflect.Float64:
		return rv.Float()
	case reflect.Complex64, reflect.Complex128:
		return real(rv.Complex())
	case reflect.String:
		if f, err := strconv.ParseFloat(rv.String(), 64); err == nil {
			return f
		}
	}

	return 0
}
