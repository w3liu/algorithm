package fibonacci

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	for i := 0; i < 10; i++ {
		fibn := Fibonacci(i)
		fmt.Println(fibn)
	}
}

func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	fib1 := 0
	fib2 := 1
	fibn := 0
	for i := 2; i <= n; i++ {
		fibn = fib1 + fib2
		fib1 = fib2
		fib2 = fibn
	}
	return fibn
}
