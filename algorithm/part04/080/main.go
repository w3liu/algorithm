package main

import "fmt"

func main() {
	var res = generateParenthesis(1)
	fmt.Println(res)
}

var result []string

func generateParenthesis(n int) []string {
	result = make([]string, 0)
	if n == 0 {
		return result
	}
	combine("", n, n)
	return result
}

func combine(str string, left, right int) {
	if left == 0 && right == 0 {
		result = append(result, str)
		return
	}
	if left == right {
		combine(str+"(", left-1, right)
	} else {
		if left > 0 {
			combine(str+"(", left-1, right)
		}
		combine(str+")", left, right-1)
	}
}
