package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}

// one pass hash table O(n)
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if _, ok := m[target-num]; ok {
			fmt.Println("ok :", ok)
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
