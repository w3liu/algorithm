package main

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	arr := []int{1, 2, 3}
	arr = arr[3:]
	fmt.Println(len(arr))
}
