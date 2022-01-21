package main

import "fmt"

func main() {
	var r = numSquares(12)
	fmt.Println(r)
}

func numSquares(n int) int {
	for n%4 == 0 {
		n = n / 4
	}
	if n%8 == 7 {
		return 4
	}
	for i := 0; i*i <= n; i++ {
		if i*i == n {
			return 1
		}
	}
	for i := 0; i*i < n; i++ {
		for j := 0; j*j < n; j++ {
			if i*i+j*j == n {
				return 2
			}
		}
	}
	return 3
}
