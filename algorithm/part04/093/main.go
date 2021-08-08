package main

import (
	"fmt"
	"math"
)

func main() {
	var grid = [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	var result = minPathSum(grid)
	fmt.Println(result)
}

func minPathSum(grid [][]int) int {
	var m, n = len(grid), len(grid[0])
	var dp = make([][]int, len(grid))
	dp[0] = make([]int, len(grid[0]))
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i] = make([]int, len(grid[0]))
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for x := 1; x < m; x++ {
		for y := 1; y < n; y++ {
			dp[x][y] = int(math.Min(float64(dp[x-1][y]), float64(dp[x][y-1]))) + grid[x][y]
		}
	}
	return dp[m-1][n-1]
}
