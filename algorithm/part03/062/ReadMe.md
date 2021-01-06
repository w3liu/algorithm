## Algorithm
### 1. 题目
```
数组元素平方处理
```
### 2. 题目描述
```
给你一个有序整数数组，数组中的数可以是正数、负数、零，请实现一个函数，这个函数返回一个整数：返回这个数组所有数的平方值中有多少种不同的取值。

举例：

nums = {-1,1,1,1} 你应该返回1，因为这个数组所有数的平方取值都是1，只有一种取值
nums = {-1,0,1,2,3} 你应该返回4，因为nums数组所有元素的平方值一共4种取值：1,0,4,9
```

### 3. 解答：
```golang
func calc(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var result, i, j int
	j = len(nums) - 1
	for i <= j {
		x := math.Abs(float64(nums[i]))
		y := math.Abs(float64(nums[j]))
		if x > y {
			result++
			for i <= j && x == math.Abs(float64(nums[i])) {
				i++
			}
		} else if x < y {
			result++
			for i <= j && y == math.Abs(float64(nums[j])) {
				j--
			}
		} else {
			result++
			for i <= j && x == math.Abs(float64(nums[i])) {
				i++
			}
			for i <= j && y == math.Abs(float64(nums[j])) {
				j--
			}
		}
	}
	return result
}
```
### 4. 说明
1. 采用双指针，一前，一后；
2. 两数的绝对值相等，那么他们的平方也相等；