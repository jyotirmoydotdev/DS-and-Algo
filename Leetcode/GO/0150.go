package main

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	var stack []int
	for _, val := range tokens {
		switch val {
		case "+":
			stack = append(stack, pop(&stack)+pop(&stack))
		case "-":
			a, b := pop(&stack), pop(&stack)
			stack = append(stack, b-a)
		case "*":
			stack = append(stack, pop(&stack)*pop(&stack))
		case "/":
			a, b := pop(&stack), pop(&stack)
			stack = append(stack, b/a)
		default:
			i, _ := strconv.Atoi(val)
			stack = append(stack, i)
		}
	}
	return stack[0]
}
func pop(stack *[]int) int {
	top := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return top
}

// func main() {
// 	slice := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
// 	fmt.Println(evalRPN(slice))
// }
