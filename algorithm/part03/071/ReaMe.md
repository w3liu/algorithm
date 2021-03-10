## Algorithm
### 1. 题目
```
8皇后问题
```
### 2. 题目描述
```
我们有一个 8x8 的棋盘，希望往里放 8 个棋子（皇后），每个棋子所在的行、列、对角线都不能有另一个棋子。
```

### 3. 解答：
```golang
var result = make([]int, 8)

var found bool

func cal8queens(row int) {
	if row == 8 {
		printQueens(result)
		found = true
		return
	}
	for j := 0; j < 8; j++ {
		if isOk(row, j) {
			result[row] = j
			cal8queens(row + 1)
			if found {
				return
			}
		}
	}
}

func isOk(row, column int) bool {
	left := column - 1
	right := column + 1
	for i := row - 1; i >= 0; i-- {
		if result[i] == column {
			return false
		}
		if left >= 0 && result[i] == left {
			return false
		}
		if right < 8 && result[i] == right {
			return false
		}
		left--
		right++
	}
	return true
}

func printQueens(result []int) {
	for i := 0; i < len(result); i++ {
		s := ""
		for j := 0; j < 8; j++ {
			if result[i] == j {
				s = fmt.Sprintf("%s %s", s, "*")
			} else {
				s = fmt.Sprintf("%s %s", s, "-")
			}
		}
		fmt.Println(s)
	}
	fmt.Println("-------------------")
}
```
### 4. 说明
1. isOk函数就是回溯算法的核心部分，通过遍历之前的点来判断当前点是否满足条件；
2. 若果当前点是OK的，则进行下一步递归调用；