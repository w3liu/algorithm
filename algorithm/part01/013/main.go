package main

import "fmt"

func main() {
	arr := relativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6})
	fmt.Println(arr)
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	indexMap := make(map[int]int)
	for i, v := range arr2 {
		indexMap[v] = i
	}
	for i := 0; i < len(arr1); i++ {
		for j := i + 1; j < len(arr1); j++ {
			vi, oki := indexMap[arr1[i]]
			vj, okj := indexMap[arr1[j]]
			if !oki || !okj {
				vi = arr1[i]
				vj = arr1[j]
			}
			if vi > vj {
				arr1[i], arr1[j] = arr1[j], arr1[i]
			}
		}
	}
	return arr1
}
