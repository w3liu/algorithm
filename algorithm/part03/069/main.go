package main

import "fmt"

func main() {
	var nums = []int{1, 3, 2, 1}
	result := majorityElement(nums)
	fmt.Println(result)
}

func majorityElement(nums []int) int {
	var candidate = nums[0]
	var cnt = 1
	for i := 1; i < len(nums); i++ {
		if candidate == nums[i] {
			cnt++
		} else {
			if cnt > 1 {
				cnt--
			} else {
				candidate = nums[i]
				cnt = 1
			}
		}
	}
	return candidate
}
