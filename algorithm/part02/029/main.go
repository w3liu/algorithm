package main

func isMatch(s string, p string) bool {
	var slen = len(s)
	var plen = len(p)
	var dp = make([][]bool, slen+1)
	for i := 0; i <= slen; i++ {
		dp[i] = make([]bool, plen+1)
	}
	for i := 0; i <= slen; i++ {
		dp[i][0] = false
	}
	dp[0][0] = true
	for i := 0; i <= slen; i++ {
		for j := 1; j <= plen; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2] || (i > 0 && (s[i-1] == p[j-2] || p[j-2] == '.') && dp[i-1][j])
			} else {
				dp[i][j] = i > 0 && dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '.')
			}
		}
	}
	return dp[slen][plen]
}
