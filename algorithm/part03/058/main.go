package main

import "fmt"

func main() {
	var arr = []int{3, 2, 2, 3}
	l := removeElement(arr, 3)
	fmt.Println(l, arr)
}

func removeElement(nums []int, val int) int {
	var i int
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
