package main

import (
	"fmt"

	"github.com/pkg/profile"
)

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	defer profile.Start(profile.ProfilePath(".")).Stop()

	fmt.Println("start")
	for i := 0; i < 1000; i++ {
		fib(30)
	}
}