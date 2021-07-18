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
	var max int
	var dp = make([]int, len(nums), len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp[0] = nums[0]
			max = nums[0]
		} else {
			temp := dp[i-1] + nums[i]
			if temp > nums[i] {
				dp[i] = temp
			} else {
				dp[i] = nums[i]
			}
			if dp[i] > max {
				max = dp[i]
			}
		}
	}
	return max
}

func maxSubArray1(nums []int) int {
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
