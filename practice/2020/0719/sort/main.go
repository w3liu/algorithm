package main

import "fmt"

func main() {
	arr := []int{1, 3, 2, 8, 7, 9, 4, 6, 5}
	arr = mSort(arr)
	printArr(arr)
}

func printArr(arr []int) {
	for _, item := range arr {
		fmt.Println(item)
	}
}

// 冒泡排序
func bubble(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		var flag bool
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return arr
}

// 插入排序
func insert(arr []int) []int {
	if len(arr) <= 0 {
		return arr
	}
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] < arr[j] {
				temp := arr[i]
				for k := i; k > j; k-- {
					arr[k] = arr[k-1]
				}
				arr[j] = temp
				break
			}
		}
	}
	return arr
}

// 选择排序
func chose(arr []int) []int {
	var temp, index int
	for i := 0; i < len(arr); i++ {
		index = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[index] {
				index = j
			}
		}
		if index > i {
			temp = arr[index]
			arr[index] = arr[i]
			arr[i] = temp
		}
	}
	return arr
}

// 归并排序
func mSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mSort(arr[:mid])
	right := mSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	l, r := 0, 0
	temp := make([]int, 0, len(left)+len(right))
	for l < len(left) && r < len(right) {
		if left[l] > right[r] {
			temp = append(temp, right[r])
			r++
		} else {
			temp = append(temp, left[l])
			l++
		}
	}
	temp = append(temp, left[l:]...)
	temp = append(temp, right[r:]...)
	return temp
}
