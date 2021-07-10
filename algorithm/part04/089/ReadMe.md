## Algorithm
### 1. 题目
```
49. 字母异位词分组
```
### 2. 题目描述
```
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

示例:

输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
```

### 3. 解答：
```golang
func groupAnagrams(strs []string) [][]string {
	var cm  = make(map[[26]int][]string)
	for _, str := range strs {
		ele := [26]int{}
		for _, s := range str {
			ele[s-'a']++
		}
		cm[ele] = append(cm[ele], str)
	}
	var result = make([][]string, 0, len(strs))
	for _, v := range cm {
		result = append(result, v)
	}
	return result
}
```
### 4. 说明
用长度为26位的数组作为map的key，遍历strs数组，给key赋值，再将字符串放入map中，最后遍历map将字符串放入二维数组。