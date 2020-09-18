package main

import (
	"fmt"
	"strings"
)

func main() {
	r := titleToNumber("ZY")
	fmt.Println("r", r)
}

func titleToNumber1(s string) int {
	var chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var base = 26
	var result int
	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		r := runes[i]
		idx := strings.IndexRune(chars, r)
		result = result + (idx+1)*exp(base, len(runes)-i-1)
	}
	return result
}

func titleToNumber(s string) int {
	var result int
	for _, r := range []rune(s) {
		num := r - 'A' + 1
		result = result*26 + int(num)
	}
	return result
}

func exp(base, n int) int {
	var result int
	for i := 0; i <= n; i++ {
		if i == 0 {
			result = 1
		} else {
			result = result * base
		}
	}
	return result
}
