package main

import (
	"fmt"
	"math"
)

func main() {
	var matrix = [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}

	var result = maximalSquare(matrix)
	fmt.Println(result)
}

func maximalSquare(matrix [][]byte) int {
	var maxSide int
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return maxSide
	}
	var rows = len(matrix)
	var cols = len(matrix[0])
	var dp = make([][]int, 0, rows)
	for i := 0; i < rows; i++ {
		dp = append(dp, make([]int, cols))
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = int(math.Min(math.Min(float64(dp[i-1][j]),
						float64(dp[i-1][j-1])),
						float64(dp[i][j-1]))) + 1
				}
			} else {
				dp[i][j] = 0
			}
			maxSide = int(math.Max(float64(maxSide), float64(dp[i][j])))
		}
	}
	return maxSide * maxSide
}
