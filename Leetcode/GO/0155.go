// 155. Min Stack
package main 
type MinStack struct {
	stack    []int // To store the elements
	minStack []int // To store the minimum values
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val) // Add the element to the main stack

	// Update the minStack with the new minimum value if necessary
	if len(this.minStack) == 0 || val <= this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack) Pop() {
	if this.isEmpty() {
		return
	}

	// Remove the element from the main stack
	popped := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]

	// Update the minStack if the popped element was the minimum
	if popped == this.minStack[len(this.minStack)-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

func (this *MinStack) Top() int {
	if this.isEmpty() {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if this.isEmpty() {
		return 0
	}
	return this.minStack[len(this.minStack)-1]
}

// ------------------------------------------------

func (this *MinStack) Len() int {
	return len(this.stack)
}

func (this *MinStack) isEmpty() bool {
	return this.Len() == 0
}
