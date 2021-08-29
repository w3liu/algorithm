package main

import "fmt"

func main() {
	arr := []int{2, 0, 2, 1, 1, 0}
	qSort(arr)
	//arr = insert(arr)
	printArr(arr)
}

func printArr(arr []int) {
	for _, item := range arr {
		fmt.Println(item)
	}
}

func insert(arr []int) []int {
	var n = len(arr)
	for i := 1; i < n; i++ {
		// 找到i的插入点
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				// 交换
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}

// 希尔排序
func shell(arr []int) []int {
	var n = len(arr)
	for gap := n / 2; gap > 0; gap /= 2 {
		for j := gap; j > 0; j = j - gap {
			if arr[j] < arr[j-gap] {
				arr[j], arr[j-gap] = arr[j-gap], arr[j]
			} else {
				break
			}
		}
	}
	return arr
}

// 快速排序
func qSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	var mid, i = arr[0], 1
	var head, tail = 0, len(arr) - 1
	for head < tail {
		if mid < arr[i] {
			arr[tail], arr[i] = arr[i], arr[tail]
			tail--
		} else {
			arr[head], arr[i] = arr[i], arr[head]
			i++
			head++
		}
	}
	qSort(arr[:head])
	qSort(arr[head+1:])
}
