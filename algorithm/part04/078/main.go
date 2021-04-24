package main

import (
	"fmt"
)

func main() {
	var results = letterCombinations("23")
	fmt.Println(results)
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var phone = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var queue = []string{""}
	for _, digit := range digits {
		for _, _ = range queue {
			temp := queue[0]
			queue = queue[1:]
			for _, letter := range phone[digit-50] {
				queue = append(queue, temp+string(letter))
			}
		}
	}
	return queue
}
