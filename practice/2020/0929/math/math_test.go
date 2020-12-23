package main

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestPow(t *testing.T) {
	begin := time.Now()
	result := math.Pow(4, 4)
	end := time.Now()
	span := end.Sub(begin)
	fmt.Println("span", span, "result", result)
}

func TestExp(t *testing.T) {
	r := exp(3)
	fmt.Println(r)
}

func exp(x int) int {
	return 1 << x
}
