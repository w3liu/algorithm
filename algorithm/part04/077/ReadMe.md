## Algorithm
### 1. 题目
```
15. 三数之和
```
### 2. 题目描述
```
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

 

示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
示例 2：

输入：nums = []
输出：[]
示例 3：

输入：nums = [0]
输出：[]
 

提示：

0 <= nums.length <= 3000
-105 <= nums[i] <= 105

```

### 3. 解答：
```golang
func threeSum(nums []int) [][]int {
	var n = len(nums)
	var results = make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := n - 1
		for l < r {
			res := nums[i] + nums[l] + nums[r]
			if res == 0 {
				result := []int{nums[i], nums[l], nums[r]}
				results = append(results, result)
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if res > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return results
}
```
### 4. 说明
* 首先对数组进行排序
* 对数组进行遍历
* 定义两个指针，一个从i+1开始为左指针，一个从n-1开始为右指针
* 判断三个数相加的值，分为三种情况，一种是等于0，一种是大于0，一种是小于0
* 然后针对三种情况分别移动指针
* 最后返回结果数组