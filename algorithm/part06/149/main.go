package main

import (
	"fmt"
	"math"
)

func main() {
	result := myAtoi("")
	fmt.Println(result)
}

func myAtoi(s string) int {
	var result int32
	var neg int
	var index int
	for ; index < len(s); index++ {
		if s[index] != ' ' {
			break
		}
	}
	if index >= len(s) {
		return 0
	}
	if s[index] == '-' {
		neg = -1
		index++
	} else if s[index] == '+' {
		neg = 1
		index++
	} else {
		neg = 1
	}
	for i := index; i < len(s); i++ {
		num := s[i]
		last := result
		if num >= '0' && num <= '9' {
			result = result*10 + (int32(num) - '0')
		} else {
			break
		}
		if result/10 != last {
			if neg == -1 {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		}
	}
	return neg * int(result)
}
