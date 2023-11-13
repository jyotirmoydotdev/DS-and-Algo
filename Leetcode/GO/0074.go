package main

import "fmt"

func main() {
	// target := 3
	slice := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	fmt.Println(searchMatrix(slice, 3))
}

func searchMatrix(matrix [][]int, target int) bool {
	ROWS := len(matrix)
	COLUMNS := len(matrix[0])

	top := 0
	bot := ROWS - 1

	for top <= bot {
		row := (top + bot) / 2
		if target > matrix[row][len(matrix[row])-1] {
			top = row + 1
		} else if target < matrix[row][0] {
			bot = row - 1
		} else {
			break
		}
	}

	if !(top <= bot) {
		return false
	}

	row := (top + bot) / 2
	left := 0
	right := COLUMNS - 1
	for left <= right {
		middle := (left + right) / 2
		if matrix[row][middle] == target {
			return true
		} else if matrix[row][middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return false

}
