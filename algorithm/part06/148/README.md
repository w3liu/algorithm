## Algorithm
### 1. 题目
```
7. 整数反转
```

### 2. 题目描述
```
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号）。


示例 1：

输入：x = 123
输出：321
示例 2：

输入：x = -123
输出：-321
示例 3：

输入：x = 120
输出：21
示例 4：

输入：x = 0
输出：0
```

### 3. 解答：
```golang
func reverse(x int) int {
	var result int
	for x != 0 {
		tmp := x % 10
		result = result*10 + tmp
		if result >= math.MaxInt32 || result <= math.MinInt32 {
			return 0
		}
		x = x / 10
	}
	return result
}
```
### 4. 说明
取模获得个位数，整除10移除个位数。