// 155 Min Stack

package main

// func main() {
// 	// Create a new MinStack instance
// 	stack := Constructor()

// 	// Push some elements onto the stack
// 	stack.Push(5)

// 	// // check the last min
// 	// fmt.Println("Last Min after 1 push: ", stack.getLastmin())

// 	stack.Push(2)
// 	stack.Push(7)
// 	stack.Push(1)

// 	// Get and print the top element
// 	// fmt.Println("Top element:", stack.Top()) // Output: 1

// 	// // Get and print the minimum element
// 	// fmt.Println("Minimum element:", stack.GetMin()) // Output: 1

// 	// // Pop an element from the stack
// 	// stack.Pop()

// 	// // Get and print the updated top element
// 	// fmt.Println("Top element after pop:", stack.Top()) // Output: 7

// 	// // Get and print the updated minimum element
// 	// fmt.Println("Minimum element after pop:", stack.GetMin()) // Output: 2

// }

type MinStack struct {
	top *StackNode
	min int
}
type StackNode struct {
	data    int
	next    *StackNode
	lastmin int
}

var mystack MinStack = MinStack{top: nil}
var newTop *StackNode

func Constructor() MinStack {
	return mystack
}

func (this *MinStack) Push(val int) {
	if this.top == nil {
		newTop = &StackNode{data: val, next: this.top}
	} else {
		newTop = &StackNode{data: val, next: this.top, lastmin: this.min}
	}
	this.top = newTop
	if this.top.data < this.min {
		this.min = this.top.data
	}
}

func (this *MinStack) Pop() {
	if this.top.next == nil {
		this.top = nil
		return
	}
	this.min = this.top.lastmin
	*this.top = *this.top.next
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
