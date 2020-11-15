package main

import "fmt"

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	result := findKthLargest(nums, k)
	fmt.Println(result)
}

func findKthLargest(nums []int, k int) int {
	return qSort(nums, k)
}

func qSort(nums []int, k int) int {
	var index int
	for i := 0; i < len(nums); i++ {
		index = i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[index] {
				index = j
			}
		}
		if index > i {
			nums[index], nums[i] = nums[i], nums[index]
		}
		if i == k-1 {
			return nums[i]
		}
	}
	return 0
}
