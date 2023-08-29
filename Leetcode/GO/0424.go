package main

// func main() {
// 	s := "AABABBA"
// 	fmt.Println("characterReplacement:", characterReplacement(s, 1))
// }

func characterReplacement(s string, k int) int {
	count := make(map[byte]int)
	maxF := 0
	left := 0
	result := 0
	for right, _ := range s {
		count[s[right]] = 1 + count[s[right]]
		maxF = max(maxF, count[s[right]])

		if (right-left+1)-maxF > k {
			count[s[left]] -= 1
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
