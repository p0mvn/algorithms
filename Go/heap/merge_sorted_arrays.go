package main

import (
	"container/heap"
)

type Item struct {
	arrayIdx int
	posIdx int
	num int
}

type MinHeap []*Item

func (mh MinHeap) Len() int { 
	return len(mh)
}

func (mh MinHeap) Less(i, j int) bool { 
	return mh[i].num < mh[j].num
}

func (mh MinHeap) Swap(i, j int) { 
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *MinHeap) Push(i interface{}) { 
	*mh = append(*mh, i.(*Item))
}

func (mh *MinHeap) Pop() interface{} { 
	length := len(*mh)
	res := (*mh)[length - 1]
	*mh = (*mh)[:length - 1]
	return res
}

// Problem: write a function that takes in a non-empty list of non-empty sorted arrays of integers and returns a merged list of all those arrays.
//
// Solution: use min heap to keep track of the smallest element in each array.
// O(nlog(k) + k) time | O(n + k) space
func MergeSortedArrays(arrays [][]int) []int {
	
	mh := make(MinHeap, len(arrays))
	for i, arr := range arrays {
		mh[i] = &Item{
			i,
			0,
			arr[0],
		}
	}
	
	heap.Init(&mh)
	
	result := make([]int, 0)
	
	for mh.Len() > 0 {
		item := heap.Pop(&mh).(*Item)
		result = append(result, item.num)
		
		curArray := arrays[item.arrayIdx]
		if item.posIdx + 1 < len(curArray) {
			// update previous to new values
			item.posIdx = item.posIdx + 1
			item.num = curArray[item.posIdx]
			heap.Push(&mh, item)
		}
	}
	
	return result
}
