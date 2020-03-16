package main

import "fmt"

func main() {
	arr := generate(5)
	fmt.Println(arr)
}

func generate(numRows int) [][]int {
	arr := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		len := i + 1
		arr[i] = make([]int, len)
		for j := 0; j < len; j++ {
			if j == 0 || j == len-1 {
				arr[i][j] = 1
			} else {
				arr[i][j] = arr[i-1][j-1] + arr[i-1][j]
			}
		}
	}
	return arr
}
