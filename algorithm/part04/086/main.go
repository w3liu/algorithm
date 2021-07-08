package main

import (
	"fmt"
	"math"
)

func main() {
	var height = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	var result = trap(height)
	fmt.Println(result)
}

func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	var ans int
	var lMax = make([]int, len(height))
	var rMax = make([]int, len(height))
	lMax[0], rMax[len(height)-1] = height[0], height[len(height)-1]
	for i := 1; i < len(height)-1; i++ {
		lMax[i] = max(height[i], lMax[i-1])
	}
	for i := len(height) - 2; i > 0; i-- {
		rMax[i] = max(height[i], rMax[i+1])
	}
	for i := 1; i < len(height)-1; i++ {
		ans = ans + min(lMax[i], rMax[i]) - height[i]
	}
	return ans
}

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func min(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}
