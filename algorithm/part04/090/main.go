package main

import (
	"fmt"
	"math"
)

func main() {
	var nums = []int{-1}
	var result = maxSubArray(nums)
	fmt.Println(result)
}
func maxSubArray(nums []int) int {
	var max = math.MinInt32
	for i := 0; i < len(nums); i++ {
		temp := nums[i]
		if temp > max {
			max = temp
		}
		for j := i + 1; j < len(nums); j++ {
			temp = temp + nums[j]
			if temp > max {
				max = temp
			}
		}
	}
	return max
}
