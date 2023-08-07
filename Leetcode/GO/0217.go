package main

import (
	"sort"
)

// func main() {
// 	nums := []int{1, 2, 3, 1, 5, 6}
// 	nums2 := []int{1, 2, 3, 4, 5, 6}
// 	fmt.Println(containsDuplicate(nums))
// 	fmt.Println(containsDuplicate(nums2))
// 	fmt.Println(containsDuplicate1(nums))
// 	fmt.Println(containsDuplicate1(nums2))
// }

// usign map
func containsDuplicate(nums []int) bool {
	nums_map := map[int]int{}
	for _, n := range nums {
		if _, ok := nums_map[n]; !ok {
			nums_map[n] = 1
		} else {
			return true
		}
	}
	return false
}

// using sort
func containsDuplicate1(nums []int) bool {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}

	return false
}
