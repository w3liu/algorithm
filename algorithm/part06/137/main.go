package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	count := findTargetSumWays(nums, target)
	fmt.Println("count", count)
}

func findTargetSumWays(nums []int, target int) int {
	var count int
	var sumfn func(index, sum int)
	sumfn = func(index, sum int) {
		if index == len(nums) {
			if sum == target {
				count++
			}
			return
		}
		sumfn(index+1, sum+nums[index])
		sumfn(index+1, sum-nums[index])
	}
	sumfn(0, 0)
	return count
}
