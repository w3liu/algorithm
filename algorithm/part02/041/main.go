package main

import "fmt"

func main() {
	n := fib(30)
	fmt.Println(n)
}

func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	return fib(N-1) + fib(N-2)
}
