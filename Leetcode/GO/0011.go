package main

import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println("maxArea:", maxArea(height))
}

func maxArea(height []int) int {
	left, right := 0, (len(height) - 1)
	result := 0

	for left < right {
		area := min(height[left], height[right]) * (right - left)
		if result < area {
			result = area
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
