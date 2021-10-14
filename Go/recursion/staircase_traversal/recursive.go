package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func StaircaseTraversal(height int, maxSteps int) int {
	if height <= 1 {
		return 1
	}
	
	total := 0
	for i := 1; i < min(maxSteps, height) + 1; i++ {
		total += StaircaseTraversal(height - i, maxSteps)
	}
	return total
}
