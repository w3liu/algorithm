package main

import "fmt"

func main() {
	piles := []int{2, 7, 9, 4, 4}
	result := stoneGameII(piles)
	fmt.Println("result", result)
}

var resultCache map[string]int
var sumCache []int
var pilesLen int

func stoneGameII(piles []int) int {
	pilesLen = len(piles)
	resultCache = make(map[string]int)
	sumCache = make([]int, len(piles)+1)
	for i := pilesLen - 1; i >= 0; i-- {
		sumCache[i] = sumCache[i+1] + piles[i]
	}
	return compute(0, 1)
}

func compute(i, m int) int {
	println(i, m)
	key := fmt.Sprintf("%d_%d", i, m)
	if _, ok := resultCache[key]; ok {
		return resultCache[key]
	}
	if i > pilesLen {
		return 0
	}
	if pilesLen <= i+2*m {
		return sumCache[i]
	}
	result := 0
	for x := 1; x < m*2+1; x++ {
		t := m
		if x > m {
			t = x
		}
		temp := sumCache[i] - compute(i+x, t)
		if result < temp {
			result = temp
		}
	}
	resultCache[key] = result
	return result
}
