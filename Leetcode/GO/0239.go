package main

// import (
// 	"fmt"
// )

// func main(){
// 	nums := []int{1,1,1,1,1,4,5}
// 	k := 6
// 	fmt.Println("maxSlidingWindow:",maxSlidingWindow(nums,k))
// 	// fmt.Println(nums[:len(nums)-1])
// }

func maxSlidingWindow(nums []int, k int) []int {
	output := []int{}
	q := make([]int,0)
	l,r := 0,0
	for r < len(nums){
		for len(q) != 0 && nums[q[len(q)-1]] < nums[r]{
			q = q[:len(q)-1]
		}
		q = append(q,r)
		if l > q[0]{
			q = q[1:]
		}
		if (r + 1) >= k{
			output = append(output, nums[q[0]])
			l++
		}
		r++
	}
	return output
}