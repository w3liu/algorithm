package main

import (
	"fmt"
	"math"
)

func main() {
	var result = minDistance("horse", "ros")
	fmt.Println(result)
}

func minDistance(word1 string, word2 string) int {
	var m, n = len(word1), len(word2)
	if m == 0 || n == 0 {
		return m + n
	}
	var dp = make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for i := 0; i < n+1; i++ {
		dp[0][i] = i
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + int(math.Min(float64(dp[i-1][j]), math.Min(float64(dp[i][j-1]), float64(dp[i-1][j-1]))))
			}
		}
	}
	return dp[m][n]
}
