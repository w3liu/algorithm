## Algorithm
### 1. 题目
```
55. 跳跃游戏
```
### 2. 题目描述
```
给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个下标。

示例 1：

输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
示例 2：

输入：nums = [3,2,1,0,4]
输出：false
解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
```

### 3. 解答：
```golang
func canJump(nums []int) bool {
	var max int
	for i := 0; i < len(nums); i++ {
		if i <= max {
			if max < i+nums[i] {
				max = i + nums[i]
			}
			if max >= len(nums)-1 {
				return true
			}
		} else {
			return false
		}
	}
	return false
}
```
### 4. 说明
* 定义变量max表示最远可以到达的位置
* 遍历数组，如果i<=max，则表示i这个位置可达
* 如果可达，则判断i+nums[i]是不是大于max
* 如果不可达，直接返回false
* 如果max大于等于len(nums)-1则返回true