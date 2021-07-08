package main

import (
	"fmt"
)

func main() {
	var nums = []int{1, 2, 3}
	var result = searchRange(nums, 1)
	fmt.Println(result)
}

func searchRange(nums []int, target int) []int {
	var min, max = -1, -1
	var l, r = 0, len(nums) - 1
	for l <= r {
		m := (l + r) >> 1
		if target == nums[m] {
			min, max = m, m
			for min-1 >= 0 && nums[min-1] == target {
				min--
			}
			for max+1 <= len(nums)-1 && nums[max+1] == target {
				max++
			}
			break
		} else {
			if target < nums[m] && target >= nums[l] {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}
	return []int{min, max}
}
