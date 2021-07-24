package main

import (
	"fmt"
)

func main() {
	var nums = []int{3, 0, 8, 2, 0, 0, 1}
	var can = canJump(nums)
	fmt.Println(can)
}

func canJump(nums []int) bool {
	var max int
	for i := 0; i < len(nums); i++ {
		if i <= max {
			if max < i+nums[i] {
				max = i + nums[i]
			}
			if max >= len(nums)-1 {
				return true
			}
		} else {
			return false
		}
	}
	return false
}
