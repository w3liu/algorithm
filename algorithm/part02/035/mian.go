package main

import "fmt"

func main() {
	result := lengthOfLastWord("hello word  ")
	fmt.Println(result)
}

func lengthOfLastWord(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}
	var wl int
	for i := l - 1; i >= 0; i-- {
		if s[i] == byte(' ') {
			if wl > 0 {
				return wl
			}
		} else {
			wl++
		}
	}
	return wl
}
