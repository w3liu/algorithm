package main

import (
	"fmt"
	"math"
)

func main() {
	var nums = []int{1, 1, 2, 2}
	result := calc(nums)
	fmt.Println("result", result)
}

func calc(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var result, i, j int
	j = len(nums) - 1
	for i <= j {
		x := math.Abs(float64(nums[i]))
		y := math.Abs(float64(nums[j]))
		if x > y {
			result++
			for i <= j && x == math.Abs(float64(nums[i])) {
				i++
			}
		} else if x < y {
			result++
			for i <= j && y == math.Abs(float64(nums[j])) {
				j--
			}
		} else {
			result++
			for i <= j && x == math.Abs(float64(nums[i])) {
				i++
			}
			for i <= j && y == math.Abs(float64(nums[j])) {
				j--
			}
		}
	}
	return result
}
