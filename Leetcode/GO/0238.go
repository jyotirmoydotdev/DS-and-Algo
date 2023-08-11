// 238. Product of Array Except Self

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}

func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	pre := 1
	pos := 1
	for i := 0; i < len(nums); i++ {
		answer[i] = pre
		pre = nums[i] * pre
	}
	for i := len(nums) - 1; i >= 0; i-- {
		answer[i] *= pos
		pos = nums[i] * pos
	}
	return answer
}

// func productExceptSelf1(nums []int) []int {
// 	answer := []int{}
// 	for i := 0; i < len(nums); i++ {
// 		product := 1
// 		for j := 0; j < len(nums); j++ {
// 			if i != j {
// 				product *= nums[j]
// 			}
// 		}
// 		answer = append(answer, product)
// 	}
// 	return answer
// }
