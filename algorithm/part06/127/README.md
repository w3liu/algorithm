## Algorithm
### 1. 题目
```
338. 比特位计数
```
### 2. 题目描述
```
给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans 作为答案。
```

### 3. 解答：
```golang
func countBits(n int) []int {
	var arr = make([]int, 0, n+1)
	for i := 0; i <= n; i++ {
		arr = append(arr, calc(i))
	}
	return arr
}

func calc(n int) int {
	var cnt int
	for {
		if n%2 == 1 {
			cnt++
		}
		n = n / 2
		if n == 0 {
			break
		}
	}
	return cnt
}
```
### 4. 说明
本道题，主要是10进制到2进制的一个转换，当n%2==1的时候表示该位为1。