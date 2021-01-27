package main

import "fmt"

func main() {
	s := "pwwkew"
	result := lengthOfLongestSubstring(s)
	fmt.Println("result", result)
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var result int
	for i := 0; i < len(s)-1; i++ {
		var cnt int
	Loop:
		for j := i + 1; j < len(s); j++ {
			for k := i; k < j; k++ {
				if s[k] == s[j] {
					break Loop
				}
			}
			cnt++
		}
		if cnt > result {
			result = cnt
		}
	}
	return result + 1
}
