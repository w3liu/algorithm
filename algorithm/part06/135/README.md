## Algorithm
### 1. 题目
```
448. 找到所有数组中消失的数字
```
### 2. 题目描述
```
给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，并以数组的形式返回结果。

 

示例 1：

输入：nums = [4,3,2,7,8,2,3,1]
输出：[5,6]
示例 2：

输入：nums = [1,1]
输出：[2]
```

### 3. 解答：
```golang
func findDisappearedNumbers(nums []int) []int {
	arrs := make([]int, len(nums)+1)
	arrs[0] = 1
	for _, v := range nums {
		arrs[v] = 1
	}
	var rets []int
	for i, v := range arrs {
		if v == 0 {
			rets = append(rets, i)
		}
	}
	return rets
}
```
### 4. 说明
将nums的数字遍历存入到arrs中，其中arrs的索引值就是nums的元素值。最后遍历arrs，找出v为0的就是元素索引组合起来就是结果。