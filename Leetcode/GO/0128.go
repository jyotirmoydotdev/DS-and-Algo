// 128. Longest Consecutive Sequence
package main

// func main() {
// 	num := []int{100, 1, 2, 3, 200, 4}
// 	fmt.Println("longestConsecutive: ", longestConsecutive(num))
// }

func longestConsecutive(nums []int) int {
	set := make(map[int]bool)

	for _, num := range nums {
		set[num] = true
	}

	res := 0
	for _, num := range nums {
		if set[num-1] {
			continue
		}

		sequence := 1
		temp := num + 1

		for set[temp] {
			sequence++
			temp++
		}
		if sequence > res {
			res = sequence
		}
	}
	return res
}
