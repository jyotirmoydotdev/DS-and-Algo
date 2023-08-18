// 15. 3Sum

package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 0, 0}
	fmt.Println("threeSum:", threeSum(nums))
}

func threeSum(nums []int) [][]int {
	quickSort(nums) //NOTE:  we can also use sort.Ints(nums)
	var result [][]int
	for num1 := 0; num1 < len(nums)-2; num1++ {
		if num1 > 0 && nums[num1] == nums[num1-1] {
			continue
		}
		num2 := num1 + 1
		num3 := len(nums) - 1
		for num2 < num3 {
			sum := nums[num2] + nums[num3] + nums[num1]
			if sum == 0 {
				result = append(result, []int{nums[num1], nums[num2], nums[num3]})
				num3--
				for num2 < num3 && nums[num3] == nums[num3+1] {
					num3--
				}
			} else if sum > 0 {
				num3--
			} else {
				num2++
			}
		}
	}
	return result
}
func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]

	left := make([]int, 0)
	right := make([]int, 0)

	for i, num := range arr {
		if i == pivotIndex {
			continue
		}

		if num <= pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	quickSort(left)
	quickSort(right)

	copy(arr, append(append(left, pivot), right...))
}
