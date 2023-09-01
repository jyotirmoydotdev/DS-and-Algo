package main

func countingSort(arr []int32) []int32 {
	// Write your code here
	var maxnum int32
	for i := 0; i < len(arr); i++ {
		if maxnum < arr[i] {
			maxnum = arr[i]
		}
	}
	newarr := make([]int32, 100)
	for i := 0; i < len(arr); i++ {
		newarr[arr[i]]++
	}
	return newarr
}
