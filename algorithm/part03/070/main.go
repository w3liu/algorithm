package main

import "fmt"

func main() {
	var arr = []int{2, 7, 9, 3, 1}
	result := rob(arr)
	fmt.Println("result", result)
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	first := nums[0]
	second := max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(first+nums[i], second)
	}
	return second
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
