## Algorithm
### 1. 题目
```
238. 除自身以外数组的乘积
```
### 2. 题目描述
```
给你一个长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。

 

示例:

输入: [1,2,3,4]
输出: [24,12,8,6]
 

提示：题目数据保证数组之中任意元素的全部前缀元素和后缀（甚至是整个数组）的乘积都在 32 位整数范围内。

说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题。

进阶：
你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）

```

### 3. 解答：
```golang
func productExceptSelf(nums []int) []int {
	var length = len(nums)
	var answer = make([]int, length, length)
	answer[0] = 1
	for i := 1; i < length; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}
	var r = 1
	for i := length - 1; i >= 0; i-- {
		answer[i] = answer[i] * r
		r = r * nums[i]
	}
	return answer
}
```
### 4. 说明
定义answer数组，先存储左边的乘积，再计算右边的乘积