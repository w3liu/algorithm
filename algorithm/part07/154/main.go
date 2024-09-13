package main

import (
	"fmt"
	"math"
)

func main() {
	var nums = []int{1, 2, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	length := removeDuplicates(nums)
	fmt.Println(length)
	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {
	n := len(nums)
	slow, fast := 2, 2
	for fast < n {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func removeDuplicates1(nums []int) int {
	var count = 1
	var max = math.MaxInt
	var idx int
	var maxcnt int
	for {
		if idx == len(nums)-1 {
			break
		}
		first := nums[idx]
		if first == nums[idx+1] {
			count = count + 1
		}
		if count == 3 {
			for ; idx < len(nums)-1; idx++ {
				if first == nums[idx+1] {
					nums[idx+1] = max
					maxcnt++
				} else {
					count = 1
					break
				}
			}
			continue
		}
		if nums[idx] != nums[idx+1] {
			count = 1
		}
		idx++
	}
	idx = 0
	for {
		if idx == len(nums) {
			break
		}
		if nums[idx] == max {
			for i := idx; i < len(nums); i++ {
				if nums[i] != max {
					nums[idx], nums[i] = nums[i], nums[idx]
					break
				}
			}
		}
		idx++
	}
	return maxcnt
}
