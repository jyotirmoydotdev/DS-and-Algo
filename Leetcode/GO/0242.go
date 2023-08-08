package main

import (
	"sort"
)

// func main() {
// 	s := "anagram"
// 	t := "nagaram"
// 	println(isAnagram2(s, t))
// }

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	runeStringS := []rune(s)
	runeStringT := []rune(t)

	sort.Slice(runeStringS, func(i, j int) bool {
		return runeStringS[i] < runeStringS[j]
	})
	sort.Slice(runeStringT, func(i, j int) bool {
		return runeStringT[i] < runeStringT[j]
	})

	sortedS := string(runeStringS)
	sortedT := string(runeStringT)

	return sortedS == sortedT
}

// O(n)
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	runeStringS := []rune(s)
	runeStringT := []rune(t)

	mapS, mapT := make(map[rune]int), make(map[rune]int)

	for i := 0; i < len(runeStringS); i++ {
		mapS[runeStringS[i]]++
		mapT[runeStringT[i]]++
	}
	for key, Value := range mapS {
		if mapT[key] != Value {
			return false
		}
	}

	return true
}
