package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Pop(stack *[]string) {
	length := len(*stack)
	*stack = (*stack)[:length-1]
}

func main() {
	RunReorderList()
}
