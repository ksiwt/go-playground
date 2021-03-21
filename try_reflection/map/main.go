package main

import (
	"fmt"
	"reflect"
)

func main() {
	rv := reflect.ValueOf(map[int]string{
		10: "a",
		20: "b",
		30: "c",
	})
	iter := rv.MapRange()
	for iter.Next() {
		k, v := iter.Key(), iter.Value()
		fmt.Println(k, v)
	}
	for _, k := range rv.MapKeys() {
		v := rv.MapIndex(k)
		fmt.Println(k, v)
	}

	// mapのイテレーションは不定であるが、以下みたいな感じ。
	// 10 a
	// 20 b
	// 30 c
}
