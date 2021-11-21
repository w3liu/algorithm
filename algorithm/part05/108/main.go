package main

import "fmt"

func main() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	var result = longestConsecutive(nums)
	fmt.Println("result:", result)
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var numMap = make(map[int]bool, len(nums))
	for _, v := range nums {
		numMap[v] = true
	}
	var longest int
	for num := range numMap {
		if !numMap[num-1] {
			currNum := num
			cnt := 1
			for numMap[currNum+1] {
				currNum++
				cnt++
			}
			if longest < cnt {
				longest = cnt
			}
		}
	}
	return longest
}
