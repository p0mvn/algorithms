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
// Solution: iterate in the direction of the sunset and
// as you approach a new building, remove all shorter buildings
// from the stack.
// Runtime: O(n), Spaec: O(n)
func SunsetViews(buildings []int, direction string) []int {
	resultStack := make ([]int, 0)
	length := len(buildings)
	
	start := 0
	increment := 1
	
	if direction == "WEST" {
		start = length - 1
		increment = -1
	}
	
	for i := start; i >= 0 && i < length; i += increment {
		cur := buildings[i]
			
		for len(resultStack) > 0 && buildings[resultStack[len(resultStack) - 1]] <= cur {
			resultStack = resultStack[:len(resultStack) - 1]
		}
		
		resultStack = append(resultStack, i)
	}

	if direction == "WEST" {
		reverse(resultStack)
	}
	
	return resultStack
}

