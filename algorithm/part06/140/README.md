## Algorithm
### 1. 题目
```
560. 和为 K 的子数组
```
### 2. 题目描述
```
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的连续子数组的个数。

示例 1：

输入：nums = [1,1,1], k = 2
输出：2
示例 2：

输入：nums = [1,2,3], k = 3
输出：2
```

### 3. 解答：
```golang
func subarraySum(nums []int, k int) int {
	var ret int
	for i := 0; i < len(nums); i++ {
		var ss int
		for y := i; y < len(nums); y++ {
			ss = ss + nums[y]
			if ss == k {
				ret++
			}
		}
	}
	return ret
}
```
### 4. 说明
第一重遍历确定子串的起始值，第二重遍历是寻找完所有的子串，如果发现ss==k不用break。