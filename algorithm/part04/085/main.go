package main

func main() {

}

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
