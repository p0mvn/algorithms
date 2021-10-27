package main

// Solution: for O(1) lookup of all methods, must store min and max at every node
type Node struct {
	val, min, max int
}

type MinMaxStack struct {
	values []*Node
}

func NewMinMaxStack() *MinMaxStack {
	return &MinMaxStack {
	}
}

func (stack *MinMaxStack) Peek() int {
	length := len(stack.values)
	if length > 0 {
		return stack.values[length - 1].val
	}
	return -1
}

func (stack *MinMaxStack) Pop() int {
	length := len(stack.values)
	if length > 0 {
		val := stack.values[length - 1].val
		stack.values = stack.values[:length - 1]
		return val
	}
	return -1
}

func (stack *MinMaxStack) Push(number int) {
	length := len(stack.values)
	
	newNode := &Node{
		number,
		number,
		number,
	}
	
	if length > 0 {
		prevTop := stack.values[length - 1]
		
		if prevTop.min < newNode.min {
			newNode.min = prevTop.min
		}
		
		if prevTop.max > newNode.max {
			newNode.max = prevTop.max
		}
	}
	
	stack.values = append(stack.values, newNode)
}

func (stack *MinMaxStack) GetMin() int {
	length := len(stack.values)
	if length > 0 {
		return stack.values[length - 1].min
	}
	return -1
}

func (stack *MinMaxStack) GetMax() int {
	length := len(stack.values)
	if length > 0 {
		return stack.values[length - 1].max
	}
	return -1
}
