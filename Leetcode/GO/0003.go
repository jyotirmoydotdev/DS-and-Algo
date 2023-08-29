package main

import "fmt"

// func main() {
// 	s := "abcabcbb"
// 	fmt.Println("lengthOfLongestSubstring:", lengthOfLongestSubstring(s))
// }

func lengthOfLongestSubstring(s string) int {
	charSet := make(map[byte]bool)
	l := 0
	res := 0
	for r, _ := range s {
		var deleted byte = 00
		exist := false
		for charSet[s[r]] {
			exist = true
			delete(charSet, s[l])
			fmt.Printf("--> %v[%v]Deleted, ", l, s[l])
			deleted = s[l]
			l++
		}
		fmt.Println("\n")
		charSet[s[r]] = true
		res = max(res, r-l+1)
		fmt.Printf("%v[%v]%v = %v[%v]Deleted = [%v]True = %v = %v\n", r, s[r], exist, l-1, deleted, s[r], s, charSet)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
