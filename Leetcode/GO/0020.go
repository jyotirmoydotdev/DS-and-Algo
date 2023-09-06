// 20. Valid Parentheses
package main
import "fmt"
func main()  {
	s:="{{}}}}}}"
	fmt.Println("isValid:",isValid(s))
}
func isValid(s string) bool {
	stack := make([]byte,0)
	pairs :=  map[byte]byte{
		')' : '(',
		']' : '[',
		'}' : '{',
	}
	for _, char :=  range []byte(s){
		pair, ok :=  pairs[char]
		fmt.Println(ok)
		if !ok{
			stack = append(stack, char)
			continue
		}
		fmt.Println(len(stack))
		if len(stack) == 0{
			fmt.Println("len(stack) == 0")
			return false 
		}
		if stack[len(stack)-1] != pair{
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}