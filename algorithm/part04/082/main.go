package main

func main() {

}

func longestValidParentheses(s string) int {
	var max int
	var dp = make([]int, len(s))
	for i, c := range s {
		if c == '(' {
			dp[i] = 0
		}
		if c == ')' {
			if i >= 1 {
				if s[i-1] == '(' {
					if i >= 2 {
						dp[i] = dp[i-2] + 2
					} else {
						dp[i] = 2
					}
				}
				if s[i-1] == ')' && i-dp[i-1] > 0 {
					if s[i-dp[i-1]-1] == '(' {
						if i-dp[i-1]-2 >= 0 {
							dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
						} else {
							dp[i] = dp[i-1] + 2
						}
					}
				}
			}
		}
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}
