## Algorithm
### 1. 题目
```
75. 颜色分类
```
### 2. 题目描述
```
给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

 

示例 1：

输入：nums = [2,0,2,1,1,0]
输出：[0,0,1,1,2,2]
示例 2：

输入：nums = [2,0,1]
输出：[0,1,2]
示例 3：

输入：nums = [0]
输出：[0]
示例 4：

输入：nums = [1]
输出：[1]
```

### 3. 解答：
```golang
func qSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	var mid, i = arr[0], 1
	var head, tail = 0, len(arr) - 1
	for head < tail {
		if mid < arr[i] {
			arr[tail], arr[i] = arr[i], arr[tail]
			tail--
		} else {
			arr[head], arr[i] = arr[i], arr[head]
			i++
			head++
		}
	}
	qSort(arr[:head])
	qSort(arr[head+1:])
}
```
### 4. 说明
这道题考察的是排序算法，原地排序算法有：希尔排序、冒泡排序、插入排序、堆排序、快速排序