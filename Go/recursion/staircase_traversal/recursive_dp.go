package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func calculate(dp []int, height int, maxSteps int) int {
	if height <= 1 {
		return 1
	}
	
	if dp[height - 1] != 0 {
		return dp[height - 1]
	}
	
	total := 0
	for i := 1; i < min(maxSteps, height) + 1; i++ {
		total += calculate(dp, height - i, maxSteps)
	}
	
	dp[height - 1] = total
	
	return dp[height - 1]
}

func StaircaseTraversal(height int, maxSteps int) int {
	dp := make([]int, height)
	return calculate(dp, height, maxSteps)
}
