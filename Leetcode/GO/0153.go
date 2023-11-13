package main

//	func main() {
//		slice := []int{4, 5, 6, 1, 2, 3}
//		fmt.Println(findMin(slice))
//	}
func findMin(nums []int) int {
	res := nums[0]
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= nums[0] {
			left = mid + 1
		} else {
			if nums[mid] < res {
				res = nums[mid]
			}
			right = mid - 1
		}
	}
	return res
}
