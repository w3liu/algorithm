## Algorithm
### 1. 题目
```
78. 子集
```
### 2. 题目描述
```
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

 

示例 1：

输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
示例 2：

输入：nums = [0]
输出：[[],[0]]
```

### 3. 解答：
```golang
func subsets(nums []int) [][]int {
	var result = make([][]int, 0)
	back(0, []int{}, nums, &result)
	return result
}

// 回溯
func back(i int, tmp, nums []int, result *[][]int) {
	// result需要传递指针
	*result = append(*result, tmp)
	for j := i; j < len(nums); j++ {
		tmp = append(tmp, nums[j])
		back(j+1, tmp, nums, result)
		dst := make([]int, len(tmp)-1)
		copy(dst, tmp) // 这里需要深拷贝
		tmp = dst
	}
}
```
### 4. 说明
这道题采用回溯算法，Golang实现需要主义的是切片的引用及深拷贝。