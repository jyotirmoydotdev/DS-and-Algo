package main

// func main() {
// 	heights := []int{2, 1, 5, 6, 2, 3}
// 	fmt.Println(largestRectangleArea(heights))
// }

func largestRectangleArea(heights []int) int {
	stack := []StackValue{}
	maxArea := 0
	var start int
	for i, h := range heights {
		start = i
		for len(stack) != 0 && stack[len(stack)-1].height > h {
			index, height := stack[len(stack)-1].index, stack[len(stack)-1].height
			stack = stack[0 : len(stack)-1]
			maxArea = max(maxArea, height*(i-index))
			start = index
		}
		stack = append(stack, StackValue{start, h})
	}
	for _, h := range stack {
		maxArea = max(maxArea, h.height*(len(heights)-h.index))
	}
	return maxArea
}

type StackValue struct {
	index  int
	height int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
