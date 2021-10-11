package main

import(
	"sort"
)

type TaskIndex struct {
	Val, Index int
}

// Return the optimal assignment of tasks to each of k workers such that the tasks
// are completed as fast as possible. k representes the number of workers. Each worker can have 2 tasks.
// Greedy solution: sort by values retaining the indexes, group the longest with the shortes tasks. As 
// a result, the max concurrent time to process all jobs is minimized.
func TaskAssignment(k int, tasks []int) [][]int {
	
	taskIndex := make([]TaskIndex, len(tasks))
	
	for i, n := range tasks {
		taskIndex[i].Val = n
		taskIndex[i].Index = i
	}
	
	sort.Slice(taskIndex, func(i, j int) bool {
		return taskIndex[i].Val < taskIndex[j].Val
	})
	
	result := make([][]int, 0)
	
	for i := 0; i < len(taskIndex) / 2; i++ {
		result = append(result, []int{taskIndex[i].Index, taskIndex[len(tasks) - 1 - i].Index})
	}
	
	return result
}
