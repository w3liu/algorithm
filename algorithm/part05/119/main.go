package main

import "fmt"

func main() {
	nums := []int{1, 3, 4, 2, 2}
	num := findDuplicate(nums)
	fmt.Println(num)
}

func findDuplicate(nums []int) int {
	var slow = nums[0]
	var fast = nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	var after = 0
	for after != slow {
		slow = nums[slow]
		after = nums[after]
	}

	return slow
}
