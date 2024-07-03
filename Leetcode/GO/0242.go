package main

import (
	"fmt"
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

	mapS, mapT := make(map[uint8]int), make(map[uint8]int)

	for i := 0; i < len(s); i++ {
		mapS[s[i]]++
		mapT[t[i]]++
	}
	for key, Value := range mapS {
		if mapT[key] != Value {
			return false
		}
	}

	return true
}

func isAnagram3(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var freq [26]int

	for idx := 0; idx < len(s); idx++ {
		fmt.Printf("s[idx]-'a' => s[%v]-'a' => %v\n", idx, s[idx])
		fmt.Printf("s[idx]-'a' => s[%v]-'a' => %v\n\n", idx, t[idx])
		freq[s[idx]-'a']++
		freq[t[idx]-'a']--
	}

	for idx := 0; idx < len(freq); idx++ {
		if freq[idx] != 0 {
			return false
		}
	}

	return true
}
