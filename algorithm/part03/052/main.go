package main

import (
	"fmt"
	"math"
)

func main() {
	result := getKthMagicNumber(5)
	fmt.Println(result)
}

func getKthMagicNumber(k int) int {
	numArr := make([]int, k, k)
	p1, p2, p3 := 0, 0, 0
	numArr[0] = 1
	for i := 1; i < k; i++ {
		numArr[i] = int(math.Min(float64(numArr[p1]*3), math.Min(float64(numArr[p2]*5), float64(numArr[p3]*7))))
		if numArr[i] == numArr[p1]*3 {
			p1++
		}
		if numArr[i] == numArr[p2]*5 {
			p2++
		}
		if numArr[i] == numArr[p3]*7 {
			p3++
		}
	}
	return numArr[k-1]
}
