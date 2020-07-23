package main

import "fmt"

func main() {
	arr := []int{1, 3, 2, 8, 7, 9, 4, 6, 5}
	arr = mSort1(arr)
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

// 冒泡排序
func bubble1(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		var flag bool
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
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
func insert1(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp := arr[j]
				for k := j; k > i; k-- {
					arr[k] = arr[k-1]
				}
				arr[i] = temp
				break
			}
		}
	}
	return arr
}

// 选择排序
func chose1(arr []int) []int {
	var index int
	for i := 0; i < len(arr); i++ {
		index = i
		for j := i + 1; j < len(arr); j++ {
			if arr[index] > arr[j] {
				index = j
			}
		}
		temp := arr[index]
		arr[index] = arr[i]
		arr[i] = temp
	}
	return arr
}

func mSort1(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mSort1(arr[:mid])
	right := mSort1(arr[mid:])
	return merge(left, right)
}

func merge1(left, right []int) []int {
	var x, y int
	arr := make([]int, 0, x+y)
	for x < len(left) && y < len(right) {
		if left[x] > right[y] {
			arr = append(arr, right[y])
			y++
		} else {
			arr = append(arr, left[x])
			x++
		}
	}
	arr = append(arr, left[x:]...)
	arr = append(arr, right[y:]...)
	return arr
}
