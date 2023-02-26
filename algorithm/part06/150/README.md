## Algorithm
### 1. 题目
```
9. 回文数
```
### 2. 题目描述
```
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

例如，121 是回文，而 123 不是。


示例 1：

输入：x = 121
输出：true
示例2：

输入：x = -121
输出：false
解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3：

输入：x = 10
输出：false
解释：从右向左读, 为 01 。因此它不是一个回文数。


提示：

-2^31<= x <= 2^31- 1
```

### 3. 解答：
```
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x%10 == 0 && x != 0 {
		return false
	}
	var reversed int
	for x > reversed {
		reversed = reversed*10 + x%10
		x = x / 10
	}
	return x == reversed || x == reversed/10
}
```
### 4. 说明
首先x小于0直接返回false，然后尾数为0的且不等于0的数也返回false。
在翻转一半的数字，与剩下的x进行比较，如果x初始值是奇数，翻转的一半需要除以10再进行比较。