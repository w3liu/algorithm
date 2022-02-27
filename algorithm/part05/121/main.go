package main

import "fmt"

func main() {
	var nums = []int{10, 9, 2, 5, 3, 7, 101, 18}
	var result = lengthOfLIS(nums)
	fmt.Println(result)
}

func lengthOfLIS(nums []int) int {
	var result int
	var dp = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		result = max(dp[i], result)
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
