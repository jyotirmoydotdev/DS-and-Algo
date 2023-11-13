// 36. Valid Sudoku
package main

import "fmt"

func main() {
	var board [][]byte = [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	fmt.Println("Is Valid Sudoku: ", isValidSudoku2(board))
}

// Time Complexity O(n2)
func isValidSudoku(board [][]byte) bool {

	hashMap := make(map[string]bool)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			row := i
			column := j

			current_val := string(board[i][j])

			if current_val == "." {
				continue
			}
			_, ok1 := hashMap[current_val+"row"+string(row)]
			_, ok2 := hashMap[current_val+"col"+string(column)]
			_, ok3 := hashMap[current_val+"gri"+string(i/3)+string(j/3)]

			if ok1 || ok2 || ok3 {
				return false
			} else {
				hashMap[current_val+"row"+string(row)] = true
				hashMap[current_val+"col"+string(column)] = true
				hashMap[current_val+"gri"+string(i/3)+string(j/3)] = true
			}
			fmt.Println(hashMap)
		}
		fmt.Println("")
	}
	return true
}

// Time Complexity O(n2) but faster than above
func isValidSudoku1(board [][]byte) bool {
	rowFlags := [9][9]bool{}
	colFlags := [9][9]bool{}
	boxFlags := [9][9]bool{}
	for i, row := range board {
		for j, b := range row {
			if '.' == b {
				continue
			}
			n := b - '1'
			if rowFlags[i][n] {
				return false
			}
			rowFlags[i][n] = true
			if colFlags[j][n] {
				return false
			}
			colFlags[j][n] = true
			if boxFlags[i/3*3+j/3][n] {
				return false
			}
			boxFlags[i/3*3+j/3][n] = true
		}
	}
	return true
}

func isValidSudoku2(board [][]byte) bool {
	if !checkRow(board) {
		return false
	}
	if !checkCol(board) {
		return false
	}
	if !checkGrid(board) {
		return false
	}
	return true
}

func checkRow(board [][]byte) bool {
	fmt.Println("inside the checkrow")
	for _, col := range board {
		seen := make(map[byte]bool)
		for _, row := range col {
			if row == 46 {
				continue
			}
			if row < 48 || row > 57 {
				return false
			}
			if seen[row] {
				return false
			}
			seen[row] = true
		}
	}
	return true
}
func checkCol(board [][]byte) bool {
	fmt.Println("inside the checkcol")
	row := 0
	for row < 9 {
		seen := make(map[byte]bool)
		col := 0
		for col < 9 {
			cell := board[col][row]
			col += 1
			if cell == 46 {
				continue
			}
			if cell < 48 || cell > 57 {
				return false
			}
			if seen[cell] {
				return false
			}
			seen[cell] = true
		}
		row += 1
	}
	fmt.Println("outside the checkcol")
	return true
}
func checkGrid(b [][]byte) bool {
	fmt.Println("inside the checkgrid")
	d := [][]int{{0, 0}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
	centerX := 1
	for centerX < 9 {
		centerY := 1
		for centerY < 9 {
			seen := make(map[byte]bool)
			for _, d_v := range d {
				d_x, d_y := d_v[0], d_v[1]
				cell := b[centerX+d_x][centerY+d_y]

				if cell == 46 {
					continue
				}

				if cell < 48 || cell > 57 {
					return false
				}

				if seen[cell] {
					return false
				}
				seen[cell] = true
			}
			centerY += 3
		}
		centerX += 3
	}
	return true
}
