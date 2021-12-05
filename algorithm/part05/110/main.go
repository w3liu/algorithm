package main

import "fmt"

func main() {
	var s = "leetcode"
	var wordDict = []string{"leet", "code"}
	ret := wordBreak(s, wordDict)
	fmt.Println(ret)
}

func wordBreak(s string, wordDict []string) bool {
	var wordMap = make(map[string]bool)
	for _, v := range wordDict {
		wordMap[v] = true
	}
	var dp = make([]bool, len(s)+1, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}
