## Algorithm
### 1. 题目
```
438. 找到字符串中所有字母异位词
```
### 2. 题目描述
```
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。

 

示例 1:

输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
 示例 2:

输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
```

### 3. 解答：
```golang
func findAnagrams(s, p string) (ans []int) {
	sl := len(s)
	pl := len(p)
	ans = make([]int, 0)
	if pl > sl {
		return
	}
	var scnt, pcnt [26]int
	for i, v := range p {
		scnt[s[i]-'a']++
		pcnt[v-'a']++
	}

	if scnt == pcnt {
		ans = append(ans, 0)
	}

	for i, v := range s[:sl-pl] {
		scnt[v-'a']--
		scnt[s[i+pl]-'a']++
		if scnt == pcnt {
			ans = append(ans, i+1)
		}
	}
	return
}
```
### 4. 说明
采用滑动窗口，需要在字符串s寻找字符串p的异位词。
因为字符串p的异位词的长度一定与字符串p的长度相同，所以我们可以在字符串s中构造一个长度为与字符串p的长度相同的滑动窗口，并在滑动中维护窗口中每种字母的数量；
当窗口中每种字母的数量与字符串p中每种字母的数量相同时，则说明当前窗口为字符串p的异位词。