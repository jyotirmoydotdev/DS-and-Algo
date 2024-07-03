package main

import "fmt"

// func main() {
// 	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
// 	groupAnagrams(strs)

// }

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[[26]int][]string)
	for _, s := range strs {
		var count [26]int
		for _, c := range s {
			count[c-'a']++
		}
		anagramMap[count] = append(anagramMap[count], s)
	}
	for k, v := range anagramMap {
		fmt.Println(k, v)
	}
	result := make([][]string, len(anagramMap))
	idx := 0
	for _, v := range anagramMap {
		result[idx] = v
		idx++
	}
	return result
}
