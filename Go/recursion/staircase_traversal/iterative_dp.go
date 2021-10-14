package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func StaircaseTraversal(height int, maxSteps int) int {
	dp := make([]int, height + 1)
	dp[0] = 1
	dp[1] = 1
	
	for i := 2; i <= height; i++ {
		for j := 1; j <= min(maxSteps, i); j++ {
			dp[i] += dp[i - j]
		}
	}
	
	return dp[height]
}
