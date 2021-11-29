package main

import "fmt"

func main() {
	var nums = []int{2, 2, 1}
	var result = singleNumber(nums)
	fmt.Println(result)
}

func singleNumber(nums []int) int {
	var single int
	for _, v := range nums {
		single ^= v
	}
	return single
}
