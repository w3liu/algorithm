## Algorithm
### 1. 题目
```
14. 最长公共前缀
```
### 2. 题目描述
```
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。

```

### 3. 解答：
```golang
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var (
		x, y, l int
		v       byte
	)
	x = math.MaxInt32
	y = len(strs)
	v = byte(0)
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if len(strs[j]) < i+1 || len(strs[j]) == 0 {
				goto RETURN
			}
			if j == 0 {
				v = strs[j][i]
			}
			if strs[j][i] != v {
				goto RETURN
			}
		}
		l++
	}
RETURN:
	return strs[0][:l]
}
```
### 4. 说明
1. golang 的字符串底层就是数组，所以支持随机访问；
2. 采用双层循环，由于不知道字符串长度，所以最外层的循环为int的最大值；
3. 内层循环需要判断字符串长度如果超过限制以及字符串长度是为零则跳出循环；
4. 内层循环如果判断到有不一样的字符串出现立即跳出循环；