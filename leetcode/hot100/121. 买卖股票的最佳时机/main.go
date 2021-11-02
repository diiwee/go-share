package main

//按着代码框架书写的
func maxProfit2(prices []int) int {

	if len(prices) < 1 {
		return 0
	}

	dp := make([][]int, len(prices)) //初始化二维切片

	for i, _ := range dp {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0 //处理越界特殊情况
	dp[0][1] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}

	return dp[len(prices)-1][0]

}

//优化后
func maxProfit(prices []int) int {

	if len(prices) < 1 {
		return 0
	}

	dp_i_0 := 0
	dp_i_1 := -prices[0] - 1

	for i := 0; i < len(prices); i++ {
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, -prices[i])
	}

	return dp_i_0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
