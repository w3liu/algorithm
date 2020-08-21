package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	if len(result) == 2 {
		fmt.Println(result[0], result[1])
	}
}

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if v, ok := cache[diff]; ok {
			return []int{v, i}
		}
		cache[nums[i]] = i
	}
	return nil
}

func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
