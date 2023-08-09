package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 3, 3, 4, 7, 8, 4, 6, 9, 33, 0}, 35))
}

// one pass hash table O(n)
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if _, ok := m[target-num]; ok {
			return []int{m[target-num], i}
		}
		m[num] = i
	}
	return []int{}
}

// Brute Force O(n^2)
func twoSum2(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
