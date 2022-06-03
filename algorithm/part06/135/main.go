package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	rets := findDisappearedNumbers(nums)
	fmt.Println(rets)
}

func findDisappearedNumbers(nums []int) []int {
	arrs := make([]int, len(nums)+1)
	arrs[0] = 1
	for _, v := range nums {
		arrs[v] = 1
	}
	var rets []int
	for i, v := range arrs {
		if v == 0 {
			rets = append(rets, i)
		}
	}
	return rets
}
