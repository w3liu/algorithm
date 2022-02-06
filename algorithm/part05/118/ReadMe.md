## Algorithm
### 1. 题目
```
283. 移动零
```
### 2. 题目描述
```
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

请注意 ，必须在不复制数组的情况下原地对数组进行操作。

示例 1:

输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]
示例 2:

输入: nums = [0]
输出: [0]
```

### 3. 解答：
```golang
func moveZeroes(nums []int) {
	var n = len(nums)
	for i := n - 1; i >= 0; i-- {
		if nums[i] == 0 {
			for j := i; j < n-1; j++ {
				if nums[j+1] == 0 {
					break
				}
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
```
### 4. 说明
遍历找到所有的零，再进行冒泡排序。