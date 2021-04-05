package main

import (
	"fmt"
	"math"
)

func main() {
	s := "babad"
	r := longestPalindrome(s)
	fmt.Println(r)
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	var start, end int
	for i := 0; i < len(s); i++ {
		odd := expand(s, i, i)
		even := expand(s, i, i+1)
		max := int(math.Max(float64(odd), float64(even)))
		if max > end-start+1 {
			start = i - (max-1)/2
			end = i + max/2
		}
	}
	return s[start : end+1]
}

func expand(s string, l, r int) int {
	for l >= 0 && r < len(s) {
		if s[l] == s[r] {
			l--
			r++
		} else {
			break
		}
	}
	return r - l - 1
}
