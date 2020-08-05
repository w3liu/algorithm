package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	i := binary1(arr, 12)
	fmt.Println("i", i)
}

func binary(arr []int, v int) int {
	l, h := 0, len(arr)-1
	for l <= h {
		mid := l + (h-l)>>1
		if arr[mid] == v {
			return mid
		}
		if arr[mid] > v {
			h = mid - 1
		}
		if arr[mid] < v {
			l = mid + 1
		}
	}
	return -1
}

func binary1(arr []int, v int) int {
	return binaryIndex(arr, 0, len(arr)-1, v)
}

func binaryIndex(arr []int, l, h, v int) int {
	if l > h {
		return -1
	}
	mid := l + (h-l)>>1
	if arr[mid] == v {
		return mid
	}
	if arr[mid] < v {
		return binaryIndex(arr, l+1, h, v)
	} else {
		return binaryIndex(arr, l, h-1, v)
	}
}
