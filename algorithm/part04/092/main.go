package main

import "fmt"

func main() {
	var result = uniquePaths(2, 3)
	fmt.Println(result)
}

func uniquePaths(m int, n int) int {
	var dp = make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for x := 1; x < m; x++ {
		for y := 1; y < n; y++ {
			dp[x][y] = dp[x-1][y] + dp[x][y-1]
		}
	}
	return dp[m-1][n-1]
}
