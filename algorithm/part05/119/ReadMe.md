## Algorithm
### 1. 题目
```
287. 寻找重复数
```
### 2. 题目描述
```
给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。

假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。

你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。

示例 1：

输入：nums = [1,3,4,2,2]
输出：2
示例 2：

输入：nums = [3,1,3,4,2]
输出：3
```

### 3. 解答：
```golang
func findDuplicate(nums []int) int {
	var slow = nums[0]
	var fast = nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	var after = 0
	for after != slow {
		slow = nums[slow]
		after = nums[after]
	}

	return slow
}
```
### 4. 说明
slow = fast 时，快慢指针相遇，slow 走过的距离是初始点（0）到环状开始的点 （x） 加上 环状开始的点（x） 到相遇点（y） 这段距离，而fast走过的距离是 初始点（0）到环状开始的点（x），点（x） 到点（y），点（y）到点（x），点（x）到点（y）。

又因为fast走过的距离是low的两倍，设0到x长度为a，x到y长度为b,则有2*（a+b） = a+ b+ (y到x的距离) + b，则y到x的距离就等于0到x的距离。

所以当新的两个指针 一个从0出发，一个从相遇点y出发时，他们走到的相同的值就是环状开始的点，即x点。