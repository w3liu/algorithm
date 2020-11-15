## Algorithm
### 1. 题目
```
215. 数组中的第K个最大元素
```
### 2. 题目描述
```
在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
示例 1:

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4

```

### 3. 解答：
```golang
func findKthLargest(nums []int, k int) int {
	return qSort(nums, k)
}

func qSort(nums []int, k int) int {
	var index int
	for i := 0; i < len(nums); i++ {
		index = i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[index] {
				index = j
			}
		}
		if index > i {
			nums[index], nums[i] = nums[i], nums[index]
		}
		if i == k-1 {
			return nums[i]
		}
	}
	return 0
}
```
### 4. 说明
使用了快速排序，实现比较简单，但是时间复杂度比较高，还可以使用堆排序，实现起来要稍微复杂点。