## Algorithm
### 1. 题目
```
39. 组合总和
```
### 2. 题目描述
```
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。 
示例 1：

输入：candidates = [2,3,6,7], target = 7,
所求解集为：
[
  [7],
  [2,2,3]
]
示例 2：

输入：candidates = [2,3,5], target = 8,
所求解集为：
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
```

### 3. 解答：
```golang
func combinationSum(candidates []int, target int) (ans [][]int) {
	var comb []int
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		// 直接跳过
		dfs(target, idx+1)
		// 选择当前数
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			// 这里表示回溯
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return
}
```
### 4. 说明
这道题参考的是别人的答案
终止条件一个是candidates用完了，一个是target是0的时候
在执行dfs的时候可以选择直接跳过也可以选择当前数
选择当前数的时候，执行完毕需要回溯