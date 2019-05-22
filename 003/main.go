package main

import (
	"fmt"
	"strings"
)

func main() {
	l := lengthOfLongestSubstring("aabaab!bb")
	fmt.Println("length", l)
}

func lengthOfLongestSubstring(s string) int {
	arr := []rune(s)
	sub := ""
	longest := ""
	for _, item := range arr {
		index := strings.Index(sub, string(item))
		if index >= 0 {
			if len(sub) > len(longest) {
				longest = sub
			}
			rs := []rune(sub)
			if len(rs) > index+1 {
				sub = fmt.Sprintf("%s%s", string((rs)[index+1:]), string(item))

			} else {
				sub = string(item)
			}
		} else {
			sub = fmt.Sprintf("%s%s", sub, string(item))
		}
	}
	if len(sub) > len(longest) {
		longest = sub
	}
	return len(longest)
}
