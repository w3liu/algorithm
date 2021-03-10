## Algorithm
### 1. 题目
```
0-1 背包
```
### 2. 题目描述
```
我们有一个背包，背包总的承载重量是 Wkg。现在我们有 n 个物品，每个物品的重量不等，并且不可分割。我们现在期望选择几件物品，装载到背包中。在不超过背包所能装载重量的前提下，如何让背包中物品的总重量最大？
```

### 3. 解答：
```golang
var maxW = math.MinInt32

func put(i, cw int, nums []int, w int) {
	if i == len(nums) || cw == w {
		if cw > maxW {
			maxW = cw
		}
		return
	}
	put(i+1, cw, nums, w)
	if cw+nums[i] <= w {
		put(i+1, cw+nums[i], nums, w)
	}
}
```
### 4. 说明
1. 当前put函数是为了求背包能装下的最大容量;
2. 递归终止条件：当背包装满，或者考察完所有的物品;
3. put(i+1, cw, nums, w)和put(i+1, cw+nums[i], nums, w)两个函数对应的就是不装入背包和装入背包两种情况。 
4. if cw+nums[i] <= w {}，这一个判断就是剪枝的过程，当装入的物品已经大于背包的容量了，就不继续执行，有效减少达不到条件的流程执行。