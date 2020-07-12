## Algorithm
### 1. 题目
```
49. 丑数
```
### 2. 题目描述
```
我们把只包含因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。

 

示例:

输入: n = 10
输出: 12
解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
说明:  

1 是丑数。
n 不超过1690。
```
### 3. 解答：
```golang
func nthUglyNumber(n int) int {
	var vector = make([]int, n, n)
	vector[0] = 1
	var p2, p3, p5 int
	for i := 1; i < n; i++ {
		vector[i] = int(math.Min(math.Min(float64(vector[p2]*2), float64(vector[p3]*3)), float64(vector[p5]*5)))
		if vector[i] == vector[p2]*2 {
			p2++
		}
		if vector[i] == vector[p3]*3 {
			p3++
		}
		if vector[i] == vector[p5]*5 {
			p5++
		}
	}
	return vector[n-1]
}
```
### 4. 说明
```
根据题意， 一个丑数必然可以写为 A0 * A1 * ..... * A(n-1) * AnA0∗A1∗.....∗A(n−1)∗An, 其中 A ∈ [2, 3, 5]A∈[2,3,5] 。那么这个丑数也可以写为 (A0 * A1 * ..... * A(n-1)) * An(A0∗A1∗.....∗A(n−1))∗An, (A0 * A1 * ..... * A(n-1))(A0∗A1∗.....∗A(n−1)) 也是个丑数， 而 An ∈ [2, 3, 5]An∈[2,3,5]， 所以一个丑数乘以 2， 3， 5 之后， 一定还是一个丑数。

并且如果从丑数序列首个元素开始 *2 *3 *5 来计算， 丑数序列也是不会产生漏解的。
丑数的排列
假设当前存在 3 个数组 nums2, nums3, nums5 分别代表丑数序列从 1 开始分别乘以 2, 3, 5 的序列， 即


nums2 = {1*2, 2*2, 3*2, 4*2, 5*2, 6*2,...}
nums3 = {1*3, 2*3, 3*3, 4*3, 5*3, 6*3,...}
nums5 = {1*5, 2*5, 3*5, 4*5, 5*5, 6*5,...}
那么， 最终的丑数序列实际上就是这 3 个有序序列对的合并结果， 计算丑数序列也就是相当于 合并 3 个有序序列。
```