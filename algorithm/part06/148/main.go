package main

import (
	"fmt"
	"math"
)

func main() {
	result := reverse(1234)
	fmt.Println(result)
}

func reverse(x int) int {
	var result int
	for x != 0 {
		tmp := x % 10
		result = result*10 + tmp
		if result >= math.MaxInt32 || result <= math.MinInt32 {
			return 0
		}
		x = x / 10
	}
	return result
}
