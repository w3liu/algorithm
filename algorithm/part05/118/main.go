package main

import "fmt"

func main() {
	var nums = []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	var n = len(nums)
	for i := n - 1; i >= 0; i-- {
		if nums[i] == 0 {
			for j := i; j < n-1; j++ {
				if nums[j+1] == 0 {
					break
				}
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
