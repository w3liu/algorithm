## Algorithm
### 1. 题目
```
169. 多数元素
```
### 2. 题目描述
```
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

 

示例 1：

输入：[3,2,3]
输出：3
示例 2：

输入：[2,2,1,1,1,2,2]
输出：2

```

### 3. 解答：
```golang

func majorityElement(nums []int) int {
	var candidate = nums[0]
	var cnt = 1
	for i := 1; i < len(nums); i++ {
		if candidate == nums[i] {
			cnt++
		} else {
			if cnt > 1 {
				cnt--
			} else {
				candidate = nums[i]
				cnt = 1
			}
		}
	}
	return candidate
}
```
### 4. 说明
1. 众数必定是大于n/2的
2. 假设第一个数就是众数，遇到与自己相等的就+1
3. 遇到与自己不相等的就-1或者更换
4. 执行到最后，candidate变量必定就是众数