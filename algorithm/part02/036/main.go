package main

import (
	"fmt"
	"math"
)

func main() {
	n := nthUglyNumber(10)
	fmt.Println(n)
}

func nthUglyNumber(n int) int {
	var vector = make([]int, n, n)
	vector[0] = 1
	var p2, p3, p5 int
	for i := 1; i < n; i++ {
		vector[i] = int(math.Min(math.Min(float64(vector[p2]*2), float64(vector[p3]*3)), float64(vector[p5]*5)))
		if vector[i] == vector[p2]*2 {
			p2++
		}
		if vector[i] == vector[p3]*3 {
			p3++
		}
		if vector[i] == vector[p5]*5 {
			p5++
		}
	}
	return vector[n-1]
}
