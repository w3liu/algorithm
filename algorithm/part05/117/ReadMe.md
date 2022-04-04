## Algorithm
### 1. 题目
```
279. 完全平方数
```
### 2. 题目描述
```
给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。

给你一个整数 n ，返回和为 n 的完全平方数的 最少数量 。

完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。

 

示例 1：

输入：n = 12
输出：3 
解释：12 = 4 + 4 + 4
示例 2：

输入：n = 13
输出：2
解释：13 = 4 + 9
```

### 3. 解答：
```golang
func numSquares(n int) int {
	for n%4 == 0 {
		n = n / 4
	}
	if n%8 == 7 {
		return 4
	}
	for i := 0; i*i <= n; i++ {
		if i*i == n {
			return 1
		}
	}
	for i := 0; i*i < n; i++ {
		for j := 0; j*j < n; j++ {
			if i*i+j*j == n {
				return 2
			}
		}
	}
	return 3
}
```
### 4. 说明
1. 四平方定理：任何正整数都可以拆分成不超过4个数的平方和 ---> 答案只可能是1,2,3,4
2. 如果一个数最少可以拆成4个数的平方和，则这个数还满足 n = (4^a)*(8b+7) ---> 因此可以先看这个数是否满足上述公式，如果不满足，答案就是1,2,3了
3. 如果这个数本来就是某个数的平方，那么答案就是1，否则答案就只剩2,3了
4. 如果答案是2，即n=a^2+b^2，那么我们可以枚举a，来验证，如果验证通过则答案是2
5. 只能是3