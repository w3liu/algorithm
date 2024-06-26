## Algorithm
### 1. 题目
```
394. 字符串解码
```
### 2. 题目描述
```
给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

 

示例 1：

输入：s = "3[a]2[bc]"
输出："aaabcbc"
示例 2：

输入：s = "3[a2[c]]"
输出："accaccacc"
示例 3：

输入：s = "2[abc]3[cd]ef"
输出："abcabccdcdcdef"
示例 4：

输入：s = "abc3[cd]xyz"
输出："abccdcdcdxyz"

```

### 3. 解答：
```golang
type strStack struct {
	arr []string
}

func newStrStack() *strStack {
	return &strStack{arr: make([]string, 0)}
}

func (s *strStack) push(c string) {
	s.arr = append(s.arr, c)
}

func (s *strStack) pop() string {
	if len(s.arr) > 0 {
		c := s.arr[len(s.arr)-1]
		s.arr = s.arr[:len(s.arr)-1]
		return c
	}
	return ""
}

type intStack struct {
	arr []int
}

func newIntStack() *intStack {
	return &intStack{arr: make([]int, 0)}
}

func (s *intStack) push(c int) {
	s.arr = append(s.arr, c)
}

func (s *intStack) pop() int {
	if len(s.arr) > 0 {
		c := s.arr[len(s.arr)-1]
		s.arr = s.arr[:len(s.arr)-1]
		return c
	}
	return -1
}

func decodeString(s string) string {
	var res string
	strs := newStrStack()
	ints := newIntStack()
	var multi int
	for _, c := range s {
		if c == '[' {
			strs.push(res)
			ints.push(multi)
			multi = 0
			res = ""
		} else if c == ']' {
			var tmp string
			currMulti := ints.pop()
			for i := 0; i < currMulti; i++ {
				tmp = fmt.Sprintf("%s%s", res, tmp)
			}
			res = strs.pop() + tmp
		} else if c >= '0' && c <= '9' {
			num, _ := strconv.ParseInt(string(c), 10, 64)
			multi = 10*multi + int(num)
		} else {
			res = res + string(c)
		}
	}
	return res
}
```
### 4. 说明
模拟两个栈，存字符串和数字，遇到`[`则将临时的字符串及数字压入栈，遇到`]`则进行出栈操作。