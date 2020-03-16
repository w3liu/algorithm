package main

import "fmt"

func main() {
	index := strStr("a", "a")
	fmt.Println("index", index)
}

func strStr(haystack string, needle string) int {
	hlen := len(haystack)
	nlen := len(needle)
	if nlen == 0 {
		return 0
	}
	hs := []rune(haystack)
	for i := 0; i < hlen; i++ {
		if hlen-i >= nlen {
			if string(hs[i:nlen+i]) == needle {
				return i
			}
		} else {
			break
		}
	}
	return -1
}
