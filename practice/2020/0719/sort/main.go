package main

import "fmt"

func main() {
	arr := []int{9, 1, 3, 2, 11, 10, 8, 7, 9, 4, 4, 6, 5, 0}
	qSort(arr)
	//arr = insert(arr)
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
func choose(arr []int) []int {
	var index int
	for i := 0; i < len(arr); i++ {
		index = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[index] {
				index = j
			}
		}
		if index > i {
			arr[index], arr[i] = arr[i], arr[index]
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

// 快速排序
func qSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	mid, i := arr[0], 1
	head, tail := 0, len(arr)-1
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

// 堆排序
func heapSort(arr []int) []int {
	l := len(arr) - 1
	buildHeap(arr, l)
	for l >= 1 {
		swap(arr, 0, l)
		l = l - 1
		heapify(arr, l, 0)
	}
	return arr
}

// 建堆
func buildHeap(arr []int, n int) {
	for i := n / 2; i >= 0; i-- {
		heapify(arr, n, i)
	}
}

// 交换
func swap(arr []int, x, y int) {
	arr[x], arr[y] = arr[y], arr[x]
}

// 自顶向下堆化
func heapify(a []int, n, i int) {
	for {
		left := i*2 + 1
		right := i*2 + 2
		var maxPos = i
		if left <= n && a[i] < a[left] {
			maxPos = left
		}
		if right <= n && a[maxPos] < a[right] {
			maxPos = right
		}
		if maxPos == i {
			break
		}
		swap(a, i, maxPos)
		i = maxPos
	}
}
