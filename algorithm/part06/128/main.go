package main

import "fmt"

func main() {
	var nums = []int{1, 2}
	var results = topKFrequent(nums, 2)
	fmt.Println(results)
}

func topKFrequent(nums []int, k int) []int {
	var numMap = make(map[int]int)
	var arrMap = make(map[int][]int)
	var arr = make([]int, 0)
	for _, n := range nums {
		if v, ok := numMap[n]; ok {
			numMap[n] = v + 1
		} else {
			numMap[n] = 1
		}
	}
	for k, v := range numMap {
		if val, ok := arrMap[v]; ok {
			arrMap[v] = append(val, k)
		} else {
			arrMap[v] = []int{k}
			arr = append(arr, v)
		}
	}
	arr = insert(arr)
	n := len(arr)

	results := make([]int, 0)
	for i := 0; i < k && i < n; i++ {
		v := arr[n-i-1]
		results = append(results, arrMap[v]...)
	}
	return results[:k]
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
