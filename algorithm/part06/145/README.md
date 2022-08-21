## Algorithm
### 1. 题目
```
739. 每日温度
```
### 2. 题目描述
```
给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。

 

示例 1:

输入: temperatures = [73,74,75,71,69,72,76,73]
输出: [1,1,4,2,1,1,0,0]
示例 2:

输入: temperatures = [30,40,50,60]
输出: [1,1,1,0]
示例 3:

输入: temperatures = [30,60,90]
输出: [1,1,0]
```

### 3. 解答：
```golang
func dailyTemperatures(temperatures []int) []int {
	predictions := make([]int, len(temperatures))
	for x := 0; x < len(temperatures); x++ {
		v := 0
		for y := x + 1; y < len(temperatures); y++ {
			if temperatures[y] > temperatures[x] {
				v = y - x
				break
			}
		}
		predictions[x] = v
	}
	return predictions
}

func dailyTemperatures1(temperatures []int) []int {
	predictions := make([]int, len(temperatures))
	stack := make([]int, 0)
	for i := 0; i < len(temperatures); i++ {
		for {
			if len(stack) == 0 {
				break
			} else {
				preIndex := stack[len(stack)-1]
				if temperatures[preIndex] < temperatures[i] {
					predictions[preIndex] = i - preIndex
					stack = stack[:len(stack)-1]
				} else {
					break
				}
			}
		}
		stack = append(stack, i)
	}
	return predictions
}
```
### 4. 说明
实现了两种方法：暴力法和单调栈，其中单点栈的出栈时机是当前的温度大于栈顶元素。