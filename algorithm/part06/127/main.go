package main

import "fmt"

func main() {
	result := countBits(5)
	fmt.Println(result)
}

func countBits(n int) []int {
	var arr = make([]int, 0, n+1)
	for i := 0; i <= n; i++ {
		arr = append(arr, calc(i))
	}
	return arr
}

func calc(n int) int {
	var cnt int
	for {
		if n%2 == 1 {
			cnt++
		}
		n = n / 2
		if n == 0 {
			break
		}
	}
	return cnt
}
