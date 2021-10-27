package main

func reverse(arr []int) {
	for i, j := 0, len(arr) - 1; i < j ; i, j = i + 1,j - 1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// Problem: direction can be "EAST" or "WEST"
// For every i, buildings[i] is the height of the ith building
// Determine the indexes of the buildings that can see the sunset 
// i.e. (find all i for which there does not exist j such that building[i] <= building[j] in the direction of the sunset)
//
// Solution: keep track of the max height seen in the direction of the sunset,
// append only the indexes of the buildings that can see the sunset
// Runtime: O(n), Spaec: O(n)
func SunsetViews(buildings []int, direction string) []int {
	result := make ([]int, 0)
	length := len(buildings)
	
	start := 0
	increment := 1
	maxSoFar := -1
	
	if direction == "EAST" {
		start = length - 1
		increment = -1
	}
	
	for i := start; i >= 0 && i < length; i += increment {
			cur := buildings[i]
			
			if cur > maxSoFar {
				result = append(result, i)
				maxSoFar = cur
			}
	}

	if direction == "EAST" {
		reverse(result)
	}
	
	return result
}
