package main

import "fmt"

func main() {
	var nums = []int{1, 2, 3}
	var result = permute(nums)
	fmt.Println(result)
}

func permute(nums []int) [][]int {
	var result = make([][]int, 0)
	backtrack(&result, nums, 0)
	return result
}

func backtrack(result *[][]int, arr []int, index int) {
	if index == len(arr) {
		temp := make([]int, 0, len(arr))
		temp = append(temp, arr...)
		*result = append(*result, temp)
	}
	for i := index; i < len(arr); i++ {
		arr[i], arr[index] = arr[index], arr[i]
		backtrack(result, arr, index+1)
		arr[i], arr[index] = arr[index], arr[i]
	}
}
