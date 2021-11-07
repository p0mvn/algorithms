// https://leetcode.com/discuss/interview-question/892956/Above-Average-Subarrays-can-we-do-better-than-O2
//
// // Not tested but seems correct from the discussion above
func aboveAverageSubarrays(arr []int) [][]int {
	// sum
	sum := 0
	for _, n := range arr {
	  sum += n
	}
	
	length := len(arr)
	
	result := make([][]int, 0)
	
	for l1 := 0; l1 < len(arr); l1++ {
	  curSum := 0
	  for r1 := 0; r1 < len(arr); r1++ {
		curSum += arr[r1]
		diffSum =: sum - curSum
		
		curLen := r1 - l1 + 1
		diffLen := length - curLen
		
		if curSum / curLen > diffSum / diffLen {
		  result = append([]int{l1, r1})
		}
	  }
	}
	
	return result
  }
  