package main

import "fmt"

func main() {
	num := climbStairs(3)
	fmt.Println(num)
}

func climbStairs(n int) int {
	var dp = make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
