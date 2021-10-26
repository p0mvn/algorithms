package main

// Visualize:
// [
//   [1, 4, 7, 12, 15, 1000],
//   [2, 5, 19, 31, 32, 1001],
//   [3, 8, 24, 33, 35, 1002],
//   [40, 41, 42, 44, 45, 1003],
//   [99, 100, 103, 106, 128, 1004]
// ]

// Start in the bottom left or top right corner of the matrix and
// continue by removing rows or columns. Can't start at the top left because
// there are 2 potential paths to take
func SearchInSortedMatrix(matrix [][]int, target int) []int {
	y := len(matrix) - 1
	x := 0
	for y >= 0 && x < len(matrix[0]) {
		cur := matrix[y][x]
		if cur < target {
			x += 1
		} else if cur > target {
			y -= 1
		} else {
			return []int {y, x}
		}
	}
	
	return []int {-1, -1}
}
