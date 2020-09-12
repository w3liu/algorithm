package main

import (
	"fmt"
	"testing"
)

func TestNumber(t *testing.T) {
	s := "0123456789"
	nums := []rune(s)
	for _, num := range nums {
		fmt.Println(num)
	}
}
