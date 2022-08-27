package main

import "fmt"

func main() {
	ret := isUnique("abcdefgg")
	fmt.Println(ret)
}

func isUnique(astr string) bool {
	if len(astr) > 26 {
		return false
	}
	arr := make([]byte, 26)
	for _, v := range astr {
		arr[v-97] = 1
	}
	var cnt byte
	for _, v := range arr {
		cnt = cnt + v
	}

	return cnt == byte(len(astr))
}
