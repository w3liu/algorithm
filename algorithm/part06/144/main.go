package main

import "fmt"

func main() {
	var s = "aaa"
	ret := countSubstrings(s)
	fmt.Println(ret)
}

func countSubstrings1(s string) int {
	ret := len(s)
	for x := 2; x <= len(s); x++ {
		for y := 0; y <= len(s)-x; y++ {
			if revert(s[y : y+x]) {
				ret++
			}
		}
	}
	return ret
}

func revert(str string) bool {
	fmt.Println(str)
	mid := len(str) / 2
	for i := 0; i < mid; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func countSubstrings(s string) int {
	n := len(s)
	ret := 0
	for i := 0; i < 2*n-1; i++ {
		l, r := i/2, i/2+i%2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			ret++
		}
	}
	return ret
}
