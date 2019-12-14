package main

import "fmt"

func main() {
	nums := []int{1, 5, 9, 2, 5, 2, 3, 4, 6, 7, 8, 2, 0, 0, 1, 3, 7, 8}
	results := mergeSort(nums)
	fmt.Println(results)
}

// 冒泡排序 最好时间复杂度O(n), 最坏时间复杂度O(n^2) 平均时间复杂度O(n^2)
func bubbleSort(nums []int) []int {
	for x := 0; x < len(nums); x++ {
		flag := true
		for y := x + 1; y < len(nums); y++ {
			if nums[x] > nums[y] {
				nums[x], nums[y] = nums[y], nums[x]
				flag = false
			}
		}
		if flag {
			break
		}
	}
	return nums
}

// 插入排序 最好时间复杂度O(n) 最坏时间复杂度O(n^2) 平均时间复杂度O(n^2)
func insertionSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		temp := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if temp < nums[j] {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = temp
	}
	return nums
}

// 归并排序 最好时间复杂度O(n*log(n)) 最坏时间复杂度O(n*log(n)) 平均时间复杂度O(n*log(n))
func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2

	left := mergeSort(nums[0:mid])
	right := mergeSort(nums[mid:])

	i := 0
	j := 0

	arr := make([]int, 0, len(nums))

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			arr = append(arr, left[i])
			i++
		} else {
			arr = append(arr, right[j])
			j++
		}
	}

	if i == len(left) {
		arr = append(arr, right[j:]...)
	} else {
		arr = append(arr, left[i:]...)
	}
	return arr
}
