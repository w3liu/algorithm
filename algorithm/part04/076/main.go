package main

import (
	"fmt"
	"math"
)

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	result := maxArea(height)
	fmt.Println("result", result)
}

func maxArea(height []int) int {
	var i, max int
	var j = len(height) - 1
	for i < j {
		h := height[i]
		if height[j] < h {
			h = height[j]
		}
		area := (j - i) * h
		if area > max {
			max = area
		}
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return max
}

func maxArea1(height []int) int {
	var max int
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			area := (j - i) * int(math.Min(float64(height[i]), float64(height[j])))
			if area > max {
				max = area
			}
		}
	}
	return max
}
