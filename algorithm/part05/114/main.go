package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	arr := productExceptSelf(nums)
	fmt.Println(arr)
}

func productExceptSelf0(nums []int) []int {
	var length = len(nums)
	var arr = make([]int, length, length)
	var left = make([]int, length, length)
	var right = make([]int, length, length)
	left[0] = 1
	right[length-1] = 1
	for i := 1; i < length; i++ {
		left[i] = left[i-1] * nums[i-1]
		right[length-i-1] = right[length-i] * nums[length-i]
	}
	for i := 0; i < length; i++ {
		arr[i] = left[i] * right[i]
	}
	return arr
}

func productExceptSelf(nums []int) []int {
	var length = len(nums)
	var answer = make([]int, length, length)
	answer[0] = 1
	for i := 1; i < length; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}
	var r = 1
	for i := length - 1; i >= 0; i-- {
		answer[i] = answer[i] * r
		r = r * nums[i]
	}
	return answer
}
