## Algorithm
### 1. 题目
```
17. 电话号码的字母组合
```
### 2. 题目描述
```
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

示例 1：

输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
示例 2：

输入：digits = ""
输出：[]
示例 3：

输入：digits = "2"
输出：["a","b","c"]
 

提示：

0 <= digits.length <= 4
digits[i] 是范围 ['2', '9'] 的一个数字。

```

### 3. 解答：
```golang
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var phone = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var queue = []string{""}
	for _, digit := range digits {
		for _, _ = range queue {
			temp := queue[0]
			queue = queue[1:]
			for _, letter := range phone[digit-50] {
				queue = append(queue, temp+string(letter))
			}
		}
	}
	return queue
}
```
### 4. 说明
1. 先将输入的 digits 中第一个数字对应的每一个字母入队
2. 然后将出队的元素与第二个数字对应的每一个字母组合后入队...直到遍历到 digits 的结尾
3. 最后队列中的元素就是所求结果