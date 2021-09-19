package main

func main() {

}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	var result int
	var heights = make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				heights[j] = heights[j] + 1
			} else {
				heights[j] = 0
			}
		}
		area := largestRectangleArea(heights)
		if area > result {
			result = area
		}
	}
	return result
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
