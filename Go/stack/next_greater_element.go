package main

type Stack []int

func(s *Stack) Pop() {
	*s = (*s)[:len(*s) - 1]
}

func(s *Stack) Peek() int {
	return (*s)[len(*s) - 1]
}

func(s *Stack) Push(n int) {
	*s = append(*s, n)
}

// Problem: find the next greater element for each element in the array
//
// Solution: use a stack to contain all the indexes of the elements that
// do not have a greater element set yet.
// On each iteration, keep popping the indices from the stack
// while current value is greater than the peek of the stack and
// update the popped index from stack with the current value.
// Once find a value that is greater than or equal to current, stop iterating over stack
// and push the current index to the stack.
// Must iterate over the entire array twice to ensure that the last element finds it's greater element
// if it is a position of len(array) - 2
// Time: O(2n)
// Space: O(n)
func NextGreaterElement(array []int) []int {
	
	stack := make(Stack, 0)
	
	result := make([]int, len(array))
	for i, _ := range array {
		result[i] = -1
	}
	
	for i := 0; i < len(array) * 2; i++  {
		iMod := i % len(array)
		for len(stack) > 0 {
			lifoIdx := stack.Peek()
			if array[iMod] <= array[lifoIdx] {
				break
			}
			result[lifoIdx] = array[iMod]
			stack.Pop()
		}
		stack.Push(iMod)
	}
	return result
}
