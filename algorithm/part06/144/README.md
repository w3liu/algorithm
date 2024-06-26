## Algorithm
### 1. 题目
```
647. 回文子串
```
### 2. 题目描述
```
给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。

回文字符串 是正着读和倒过来读一样的字符串。

子字符串 是字符串中的由连续字符组成的一个序列。

具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。

 

示例 1：

输入：s = "abc"
输出：3
解释：三个回文子串: "a", "b", "c"
示例 2：

输入：s = "aaa"
输出：6
解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
```

### 3. 解答：
```golang
func countSubstrings1(s string) int {
	ret := len(s)
	for x := 2; x <= len(s); x++ {
		for y := 0; y <= len(s)-x; y++ {
			if revert(s[y : y+x]) {
				ret++
			}
		}
	}
	return ret
}

func revert(str string) bool {
	fmt.Println(str)
	mid := len(str) / 2
	for i := 0; i < mid; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func countSubstrings(s string) int {
	n := len(s)
	ret := 0
	for i := 0; i < 2*n-1; i++ {
		l, r := i/2, i/2+i%2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			ret++
		}
	}
	return ret
}
```
### 4. 说明
方法一：是枚举法时间复杂度很高，会超时。
方法二：是中心扩展法，遍历每个可能的中心点，时间复杂度相对较低。