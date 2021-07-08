## Algorithm
### 1. 题目
```
34. 在排序数组中查找元素的第一个和最后一个位置
```
### 2. 题目描述
```
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

进阶：

你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？
 

示例 1：

输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例 2：

输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：

输入：nums = [], target = 0
输出：[-1,-1]
```

### 3. 解答：
```golang
func searchRange(nums []int, target int) []int {
	var min, max = -1, -1
	var l, r = 0, len(nums) - 1
	for l <= r {
		m := (l + r) >> 1
		if target == nums[m] {
			min, max = m, m
			for min-1 >= 0 && nums[min-1] == target {
				min--
			}
			for max+1 <= len(nums)-1 && nums[max+1] == target {
				max++
			}
			break
		} else {
			if target < nums[m] && target >= nums[l] {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}
	return []int{min, max}
}
```
### 4. 说明
1. 采用二分查找能满足O(logN)的时间复杂度
2. 找到target之后，再移动左右指针确定边界