package sort

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	arr := []int{3, 4, 1, 2, 5}
	insertionSort1(arr)
	fmt.Println(arr)
}

func TestSelectSort(t *testing.T) {
	arr := []int{3, 4, 1, 2, 5}
	selectSort1(arr)
	fmt.Println(arr)
}

func TestMergeSort(t *testing.T) {
	arr := []int{3, 4, 1, 2, 5}
	ret := mergeSort1(arr)
	fmt.Println(ret)
}

func TestQuickSort(t *testing.T) {
	arr := []int{3, 4, 1, 2, 5}
	quickSort(arr)
	fmt.Println(arr)
}

func insertionSort1(arr []int) {
	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		var j = i - 1
		for ; j >= 0; j-- {
			if arr[j] > arr[i] {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = tmp
	}
}

func selectSort1(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min > i {
			arr[min], arr[i] = arr[i], arr[min]
		}
	}
}

func mergeSort1(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort1(arr[:mid])
	right := mergeSort1(arr[mid:])
	return merge1(left, right)
}

func merge1(left, right []int) []int {
	var l, r int
	var result = make([]int, 0, l+r)
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}
