package main

import "fmt"

func main() {
	var result = numTrees(3)
	fmt.Println(result)
}

func numTrees(n int) int {
	var dp = make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] = dp[i] + dp[j-1]*dp[i-j]
		}
	}
	return dp[n]
}
