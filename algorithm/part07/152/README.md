## Algorithm
### 1. 题目
```
13. 罗马数字转整数
```
### 2. 题目描述
```
罗马数字包含以下七种字符:I，V，X，L，C，D和M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做II，即为两个并列的 1 。12 写做XII，即为X+II。 27 写做XXVII, 即为XX+V+II。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做IIII，而是IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为IX。这个特殊的规则只适用于以下六种情况：

I可以放在V(5) 和X(10) 的左边，来表示 4 和 9。
X可以放在L(50) 和C(100) 的左边，来表示 40 和90。
C可以放在D(500) 和M(1000) 的左边，来表示400 和900。
给定一个罗马数字，将其转换成整数。



示例1:

输入:s = "III"
输出: 3
示例2:

输入:s = "IV"
输出: 4
示例3:

输入:s = "IX"
输出: 9
示例4:

输入:s = "LVIII"
输出: 58
解释: L = 50, V= 5, III = 3.
示例5:

输入:s = "MCMXCIV"
输出: 1994
解释: M = 1000, CM = 900, XC = 90, IV = 4.
```

### 3. 解答：
```
func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}
	var preNum = getNum(s[0])
	var total int
	for i := 1; i < len(s); i++ {
		currNum := getNum(s[i])
		if currNum > preNum {
			total -= preNum
		} else {
			total += preNum
		}
		preNum = currNum
	}
	total += preNum
	return total
}

func getNum(ch uint8) int {
	switch ch {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}
```
### 4. 说明
从左往右，左边的字符代表的数字如果小于右边则应该被减掉，反之则应该加上它。