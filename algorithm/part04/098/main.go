package main

import (
	"fmt"
)

func main() {
	var heights = []int{2, 1, 2}
	result := largestRectangleArea(heights)
	fmt.Println(result)
}

func largestRectangleArea(heights []int) int {
	var n = len(heights)
	var res = 0
	var stack = make([]int, 0)
	var temp = make([]int, n+2, n+2)
	for i := 0; i < n; i++ {
		temp[i+1] = heights[i]
	}
	heights = temp
	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 {
			tail := stack[len(stack)-1]
			if heights[i] < heights[tail] {
				stack = stack[:len(stack)-1]
				curI := tail
				curH := heights[curI]
				if len(stack) > 0 {
					tail = stack[len(stack)-1]
					leftI := tail
					rightI := i
					curW := rightI - leftI - 1
					curS := curW * curH
					if curS > res {
						res = curS
					}
				}
			} else {
				break
			}
		}
		stack = append(stack, i)
	}
	return res
}
