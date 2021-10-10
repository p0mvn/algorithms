package main

type Position struct {
	y, x int
}

type Queue []*Position

func(q *Queue) Push(n *Position) {
	*q = append(*q, n)
}

func(q *Queue) Pop() *Position {
	if len(*q) > 0 {
		toReturn := (*q)[0]
		// Ensure the removed element is garbage collected
		// If appends are not frequent, may force garbage collector to execute in Pop to optimize for memory
		// More info: https://stackoverflow.com/questions/28432658/does-go-garbage-collect-parts-of-slices
		(*q)[0] = nil
		*q = (*q)[1:]
		return toReturn
	}
	
	return nil
}

func(q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func convertToPositiveIfValid(matrix [][]int, q *Queue, y, x int) {
	if !(y >= 0 && y < len(matrix)) {
		return
	}
	
	if !(x >= 0 && x < len(matrix[y])) {
		return
	}
	
	if matrix[y][x] < 0 {
		matrix[y][x] = matrix[y][x] * -1
		q.Push(&Position{y, x})
	}
}

func breadthFirstIterate(matrix [][]int, queue Queue, iterationCount int) int {
	if queue.IsEmpty() {
		return iterationCount
	}
	
	currentLength := len(queue)
	
	for currentLength > 0 {
		pos := queue.Pop()
		convertToPositiveIfValid(matrix, &queue, pos.y - 1, pos.x)
		convertToPositiveIfValid(matrix, &queue, pos.y + 1, pos.x)
		convertToPositiveIfValid(matrix, &queue, pos.y, pos.x - 1)
		convertToPositiveIfValid(matrix, &queue, pos.y, pos.x + 1)
		
		currentLength--
	}
	
	return breadthFirstIterate(matrix, queue, iterationCount + 1)
}

// Write a function that takes in an integer matix of potentially unequal height and width and returns the minimum number of passes required to convert all
// negative integers in the matrix to positive integers.
//
// Solution: start from enqueueing positive integers and do BFS on negative integers.
// Each BFS iteration on a layer of negative numbers is one iteration
// O(w * h) time and O(w * h) space where w is the width and h is the height
func MinimumPassesOfMatrix(matrix [][]int) int {
	positiveNumQueue := make(Queue, 0)
	
	for y, row := range matrix {
		for x, elemValue := range row {
			if elemValue > 0 {
				positiveNumQueue.Push(&Position{y, x})
			}
		}
	}
	
	// Starting from -1 to account for iterating over positive numbers
	// in the first round.
	iterationCount := breadthFirstIterate(matrix, positiveNumQueue, -1)
	
	for _, row := range matrix {
		for _, elemValue := range row {
			if elemValue < 0 {
				return -1
			}
		}
	}
	
	return iterationCount
}
