## Algorithm
### 1. 题目
```
33. 搜索旋转排序数组
```
### 2. 题目描述
```
整数数组 nums 按升序排列，数组中的值 互不相同 。

在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。

给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。

 

示例 1：

输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4
示例 2：

输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1
示例 3：

输入：nums = [1], target = 0
输出：-1

你可以设计一个时间复杂度为 O(log n) 的解决方案吗？
```

### 3. 解答：
```golang
func search(nums []int, target int) int {
	var l, mid, r = 0, 0, len(nums) - 1
	for l <= r {
		mid = l + (r-l)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[l] <= nums[mid] {
			if nums[l] <= target && nums[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && nums[r] >= target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
```
### 4. 说明
* 二分法可以实现O(logN)的时间复杂度
* 二分之后一定有一边是有序的
* 然后判断搜索目标是否在有序数组中
* 如果在其中则按传统的二分法进行搜索
* 如果不在其中则切换到另一半数组再进行搜索