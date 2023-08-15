// 36. Valid Sudoku
package main

// func main() {
// 	var board [][]byte = [][]byte{
// 		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
// 		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
// 		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
// 		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
// 		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
// 		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
// 		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
// 		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
// 		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
// 	fmt.Println("Is Valid Sudoku: ", isValidSudoku(board))
// }

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
		}
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
