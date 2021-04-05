## Algorithm
### 1. 题目
```
5. 最长回文子串
```
### 2. 题目描述
```
给你一个字符串 s，找到 s 中最长的回文子串。

 

示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"
示例 3：

输入：s = "a"
输出："a"
示例 4：

输入：s = "ac"
输出："a"
```

### 3. 解答：
```golang
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	var start, end int
	for i := 0; i < len(s); i++ {
		odd := expand(s, i, i)
		even := expand(s, i, i+1)
		max := int(math.Max(float64(odd), float64(even)))
		if max > end-start+1 {
			start = i - (max-1)/2
			end = i + max/2
		}
	}
	return s[start : end+1]
}

func expand(s string, l, r int) int {
	for l >= 0 && r < len(s) {
		if s[l] == s[r] {
			l--
			r++
		} else {
			break
		}
	}
	return r - l - 1
}
```
### 4. 说明
* 中心点，有两种情况，一种是奇数的中心点，一种是偶数的中心点
* 根据中心点向两端展开，并返回展开的长度
* 取奇数与偶数展开的最大值
* 将最大值与之前的end、start进行比较
* 最后返回start与end之间的字符串