## Algorithm
### 1. 题目
```
198. 打家劫舍
```
### 2. 题目描述
```
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

 

示例 1：

输入：[1,2,3,1]
输出：4
解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
     偷窃到的最高金额 = 1 + 3 = 4 。
示例 2：

输入：[2,7,9,3,1]
输出：12
解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
     偷窃到的最高金额 = 2 + 9 + 1 = 12 。
 

提示：

0 <= nums.length <= 100
0 <= nums[i] <= 400
```

### 3. 解答：
```golang
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	var mn = make([]int, len(nums), len(nums))
	mn[0] = nums[0]
	mn[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		mn[i] = max(mn[i-2]+nums[i], mn[i-1])
	}
	return mn[len(nums)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
```
### 4. 说明
* 采用动态规划+滚动数组
* 状态转移方程：
```text
mn[i] = max(mn[i-2] + nums[i], mn[i-1])
```
* 边界条件：
```text
mn[0] = nums[0]
mn[1] = max(nums[0], nums[1])
```
* 最终答案为`mn[n-1]`，其中n为数组的长度 