package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 4}
	cnt := findUnsortedSubarray(nums)
	fmt.Println("cnt:", cnt)
}

func findUnsortedSubarray(nums []int) int {
	var cnt int
	sortedNum := make([]int, len(nums))
	_ = copy(sortedNum, nums)
	sort.Ints(sortedNum)

	var i, j = 0, len(nums) - 1

	for i <= j {
		if nums[i] == sortedNum[i] {
			cnt++
			i++
		} else if nums[j] == sortedNum[j] {
			cnt++
			j--
		} else {
			break
		}
	}
	return len(nums) - cnt
}
