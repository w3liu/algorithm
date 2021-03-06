## Algorithm
### 1. 题目
```
32. 最长有效括号
```
### 2. 题目描述
```
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。

 

示例 1：

输入：s = "(()"
输出：2
解释：最长有效括号子串是 "()"
示例 2：

输入：s = ")()())"
输出：4
解释：最长有效括号子串是 "()()"
示例 3：

输入：s = ""
输出：0

```

### 3. 解答：
```golang
func longestValidParentheses(s string) int {
	var max int
	var dp = make([]int, len(s), len(s))
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
```
### 4. 说明
1. 采用动态规划，定义数组dp，用于存放有效字符串长度；
2. 如果s[i]=')', s[i-1]='(', dp[i]=dp[i-2] + 2
3. 如果s[i]=')', s[i-1]=')', 如果s[i-dp[i-1]-1]='(', 那么dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2