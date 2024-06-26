## Algorithm
### 1. 题目
```
347. 前 K 个高频元素
```
### 2. 题目描述
```
给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。

示例 1:

输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:

输入: nums = [1], k = 1
输出: [1]
```

### 3. 解答：
```golang
func topKFrequent(nums []int, k int) []int {
	var numMap = make(map[int]int)
	var arrMap = make(map[int][]int)
	var arr = make([]int, 0)
	for _, n := range nums {
		if v, ok := numMap[n]; ok {
			numMap[n] = v + 1
		} else {
			numMap[n] = 1
		}
	}
	for k, v := range numMap {
		if val, ok := arrMap[v]; ok {
			arrMap[v] = append(val, k)
		} else {
			arrMap[v] = []int{k}
			arr = append(arr, v)
		}
	}
	arr = insert(arr)
	n := len(arr)

	results := make([]int, 0)
	for i := 0; i < k && i < n; i++ {
		v := arr[n-i-1]
		results = append(results, arrMap[v]...)
	}
	return results[:k]
}

// 插入排序
func insert(arr []int) []int {
	if len(arr) <= 0 {
		return arr
	}
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] < arr[j] {
				temp := arr[i]
				for k := i; k > j; k-- {
					arr[k] = arr[k-1]
				}
				arr[j] = temp
				break
			}
		}
	}
	return arr
}
```
### 4. 说明
采用两个map分别记录数字出现的次数以及次数对应的数组，然后再对arr进行排序。最后返回results