package main

import "fmt"

func main() {
	nums := []int{1, 1, 1}
	ret := subarraySum(nums, 2)
	fmt.Println(ret)
}

func subarraySum(nums []int, k int) int {
	var ret int
	for i := 0; i < len(nums); i++ {
		var ss int
		for y := i; y < len(nums); y++ {
			ss = ss + nums[y]
			if ss == k {
				ret++
			}
		}
	}
	return ret
}
