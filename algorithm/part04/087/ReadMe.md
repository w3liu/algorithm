## Algorithm
### 1. 题目
```
46. 全排列
```
### 2. 题目描述
```
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

 

示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
示例 2：

输入：nums = [0,1]
输出：[[0,1],[1,0]]
示例 3：

输入：nums = [1]
输出：[[1]]

```

### 3. 解答：
```golang
func permute(nums []int) [][]int {
	var result = make([][]int, 0)
	backtrack(&result, nums, 0)
	return result
}

func backtrack(result *[][]int, arr []int, index int) {
	if index == len(arr) {
		temp := make([]int, 0, len(arr))
		temp = append(temp, arr...)
		*result = append(*result, temp)
	}
	for i := index; i < len(arr); i++ {
		arr[i], arr[index] = arr[index], arr[i]
		backtrack(result, arr, index+1)
		arr[i], arr[index] = arr[index], arr[i]
	}
}
```
### 4. 说明
本题的思路就是回溯算法的具体应用：
* 递归函数backtrack的result参数用于保存结果，arr表示中间结果，index表示游标
* 游标从0开始，如果index等于arr的长度表示遍历结束将arr添加到result里
* 从游标位置开始遍历后边的每一个元素，有三个操作：
    - 交换i与index对应的元素
    - index+1继续向下迭代
    - 再次交换i与index的位置（即还原），这一步也是回溯算法的关键
