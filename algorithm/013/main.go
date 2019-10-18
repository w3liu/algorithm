package main

import "fmt"

func main() {
	arr := relativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6})
	fmt.Println(arr)
}

func relativeSortArray(arr1 []int, arr2 []int) []int {

	tempMap := make(map[int]int)
	tempArr := make([]int, 0)
	lastArr := make([]int, 0)

	for _, x := range arr1 {
		have := false
		for _, y := range arr2 {
			if x == y {
				have = true
				if v, ok := tempMap[x]; ok {
					tempMap[x] = v + 1
				} else {
					tempMap[x] = 1
				}
			}
		}
		if !have {
			lastArr = append(lastArr, x)
		}
	}
	for _, x := range arr2 {
		for i := 0; i < tempMap[x]; i++ {
			tempArr = append(tempArr, x)
		}
	}
	for i := 0; i < len(lastArr); i++ {
		for k := i + 1; k < len(lastArr); k++ {
			if lastArr[i] > lastArr[k] {
				lastArr[k], lastArr[i] = lastArr[i], lastArr[k]
			}
		}
	}
	tempArr = append(tempArr, lastArr...)
	return tempArr
}
