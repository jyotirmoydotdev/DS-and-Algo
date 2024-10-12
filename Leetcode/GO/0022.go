package main

import (
	"strings"
)

// func main() {
// 	ans := generateParenthesis(3)
// 	fmt.Println("ans:", ans)
// 	for i, e := range ans {
// 		fmt.Printf("%v = %v\n", i, e)
// 	}
// }

func generateParenthesis(n int) []string {
	var stack []string
	var res []string

	var backtrack func(int, int)
	backtrack = func(openN, closedN int) {
		if openN == n && closedN == n && openN == closedN {
			res = append(res, strings.Join(stack, ""))
			return
		}
		if openN < n {
			stack = append(stack, "(")
			backtrack(openN+1, closedN)
			Pop(&stack)
		}
		if openN > closedN {
			stack = append(stack, ")")
			backtrack(openN, closedN+1)
			Pop(&stack)
		}
	}
	backtrack(0, 0)
	return res
}
