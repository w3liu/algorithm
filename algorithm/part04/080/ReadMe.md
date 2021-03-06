## Algorithm
### 1. 题目
```
22. 括号生成
```
### 2. 题目描述
```
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

 

示例 1：

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
示例 2：

输入：n = 1
输出：["()"]
 

提示：

1 <= n <= 8

```

### 3. 解答：
```golang
func main() {
	var res = generateParenthesis(1)
	fmt.Println(res)
}

var result []string

func generateParenthesis(n int) []string {
	result = make([]string, 0)
	if n == 0 {
		return result
	}
	combine("", n, n)
	return result
}

func combine(str string, left, right int) {
	if left == 0 && right == 0 {
		result = append(result, str)
		return
	}
	if left == right {
		combine(str+"(", left-1, right)
	} else {
		if left > 0 {
			combine(str+"(", left-1, right)
		}
		combine(str+")", left, right-1)
	}
}
```
### 4. 说明
1. 剩余左括号与右括号相等则应该添加左括号
2. 如果剩余左括号小于右括号且左括号大于0，则可以添加左括号或者右括号
3. 如果剩余左括号等于0右括号大于零则继续添加右括号
4. 最后左右括号都添加完毕加入到结果集中
