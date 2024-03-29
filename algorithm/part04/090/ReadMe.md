## Algorithm
### 1. 题目
```
53. 最大子序和
```
### 2. 题目描述
```
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
示例 2：

输入：nums = [1]
输出：1
示例 3：

输入：nums = [0]
输出：0
示例 4：

输入：nums = [-1]
输出：-1
示例 5：

输入：nums = [-100000]
输出：-100000
```

### 3. 解答：
```golang
func maxSubArray(nums []int) int {
	var max int
	var dp = make([]int, len(nums), len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp[0] = nums[0]
			max = nums[0]
		} else {
			temp := dp[i-1] + nums[i]
			if temp > nums[i] {
				dp[i] = temp
			} else {
				dp[i] = nums[i]
			}
			if dp[i] > max {
				max = dp[i]
			}
		}
	}
	return max
}
```
### 4. 说明
动态规划，dp用于存放以i结尾的最大值，动态转移方程 `f(i) = max(dp[i-1] + nums[i], nums[i])`