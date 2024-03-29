## Algorithm
### 1. 题目
```
面试题 01.01. 判定字符是否唯一
```
### 2. 题目描述
```
实现一个算法，确定一个字符串 s 的所有字符是否全都不同。

示例 1：

输入: s = "leetcode"
输出: false 
示例 2：

输入: s = "abc"
输出: true
```

### 3. 解答：
```golang
func isUnique(astr string) bool {
	if len(astr) > 26 {
		return false
	}
	arr := make([]byte, 26)
	for _, v := range astr {
		arr[v-97] = 1
	}
	var cnt byte
	for _, v := range arr {
		cnt = cnt + v
	}

	return cnt == byte(len(astr))
}
```
### 4. 说明
1. 字母最多26个，因此如果字符串超过26个必然存在重复，直接返回即可。
2. 可以定义一个26位的byte数组，用于标记字母是否存在。
3. 由于a的ASCII码为97，因此遍历字符串的每个字符并减去97之后作为索引，对应的值设为1。
4. 遍历数组，统计1的个数标记为cnt，最后检查cnt是否与字符串的长度相等。