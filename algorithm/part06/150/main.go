package main

import "fmt"

func main() {
	ret := isPalindrome(12321)
	fmt.Println(ret)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x%10 == 0 && x != 0 {
		return false
	}
	var reversed int
	for x > reversed {
		reversed = reversed*10 + x%10
		x = x / 10
	}
	return x == reversed || x == reversed/10
}
