package main

import (
	"fmt"
	"testing"
)

func GetFn() func() {
	fmt.Println("outside")
	return func() {
		fmt.Println("Inside")
	}
}

func TestDefer(t *testing.T) {
	defer GetFn()()
	fmt.Println("Here")
}
