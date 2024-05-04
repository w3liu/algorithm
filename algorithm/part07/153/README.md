## Algorithm
### 1. 题目
```
16. 最接近的三数之和
```
### 2. 题目描述
```
给你一个长度为 n 的整数数组nums和 一个目标值target。请你从 nums 中选出三个整数，使它们的和与target最接近。

返回这三个数的和。

假定每组输入只存在恰好一个解。



示例 1：

输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
示例 2：

输入：nums = [0,0,0], target = 1
输出：0
```

### 3. 解答：
```
func threeSumClosest1(nums []int, target int) int {
	sort.Ints(nums)
	var result = math.MaxInt64
	var delta = math.MaxInt64
	var size = len(nums)
	for x := 0; x < size-2; x++ {
		y := x + 1
		z := size - 1
		for y < z {
			val := nums[x] + nums[y] + nums[z]
			if val == target {
				return val
			}
			sub := getAbs(val - target)
			if delta > sub {
				delta = sub
				result = val
			}
			if val > target {
				z--
			} else {
				y++
			}
		}
	}
	return result
}
```
### 4. 说明
1. 先排序
2. 遍历数组nums作为第一数字
3. 后面两个数组采用头尾双指针的方式移动并计算