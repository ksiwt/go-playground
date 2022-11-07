package main

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("hello")
	m.Run()
	fmt.Println("world")
}

func TestA(t *testing.T) {
	fmt.Println("TestA")
}

func TestB(t *testing.T) {
	fmt.Println("TestB")
}
