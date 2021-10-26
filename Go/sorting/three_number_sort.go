package main

func swap(array []int, idxA, idxB int) {
	array[idxA], array[idxB] = array[idxB], array[idxA]
}

// array contains up to 3 options of different integers
// order specifies the order of the options in array
// 
// Solution:
//
// 2 invariants:
// 1. left part of the array contains order[0]
// 2. Right part of the array contains order[2]
//
// Runtime: O(n), Space complexity: O(1)
func ThreeNumberSort(array []int, order []int) []int {
	
	nextSwapLeftIdx := 0
	nextSwapRightIdx := len(array) - 1
	
	i := 0
	for i <= nextSwapRightIdx {
		n := array[i]
		
		if n == order[0] {
			swap(array, nextSwapLeftIdx, i)
			i++
			nextSwapLeftIdx++
		} else if n == order[2] {
			swap(array, nextSwapRightIdx, i)
			nextSwapRightIdx--
		} else {
			i++
		}
	}
	
	return array
}
