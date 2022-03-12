package main

func main() {

}

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	// f[i][0]: maximum yield from holding stocks
	// f[i][1]: do not own stock, and in the frozen period of the cumulative maximum yield
	// f[i][2]: do not own stock, and not in the freeze period cumulative maximum yield
	f := make([][3]int, n)
	f[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		f[i][0] = max(f[i-1][0], f[i-1][2]-prices[i])
		f[i][1] = f[i-1][0] + prices[i]
		f[i][2] = max(f[i-1][1], f[i-1][2])
	}
	return max(f[n-1][1], f[n-1][2])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
