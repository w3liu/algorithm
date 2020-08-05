package main

import (
	"fmt"
	"testing"
)

func TestMoveBit(t *testing.T) {
	l, h := 98, 100
	mid := l + (h-l)>>1
	fmt.Println(mid)
}
