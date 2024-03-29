## Algorithm
### 1. 题目
```
581. 最短无序连续子数组
```
### 2. 题目描述
```
给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。

请你找出符合题意的 最短 子数组，并输出它的长度。

 

示例 1：

输入：nums = [2,6,4,8,10,9,15]
输出：5
解释：你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
示例 2：

输入：nums = [1,2,3,4]
输出：0
示例 3：

输入：nums = [1]
输出：0
```

### 3. 解答：
```golang
func findUnsortedSubarray(nums []int) int {
	var cnt int
	sortedNums := make([]int, len(nums))
	_ = copy(sortedNum, nums)
	sort.Ints(sortedNum)

	var i, j = 0, len(nums) - 1

	for i <= j {
		if nums[i] == sortedNum[i] {
			cnt++
			i++
		} else if nums[j] == sortedNum[j] {
			cnt++
			j--
		} else {
			break
		}
	}
	return len(nums) - cnt
}
```
### 4. 说明
先排序，然后再对比头尾有多少是相同的，再用总长度减去首尾长度之和即可。
