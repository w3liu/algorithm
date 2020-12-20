package main

import (
	"fmt"
	"math"
)

func main() {
	strs := []string{""}
	result := longestCommonPrefix(strs)
	fmt.Println("result", result)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var (
		x, y, l int
		v       byte
	)
	x = math.MaxInt32
	y = len(strs)
	v = byte(0)
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if len(strs[j]) < i+1 || len(strs[j]) == 0 {
				goto RETURN
			}
			if j == 0 {
				v = strs[j][i]
			}
			if strs[j][i] != v {
				goto RETURN
			}
		}
		l++
	}
RETURN:
	return strs[0][:l]
}
