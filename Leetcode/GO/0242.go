package main

import (
	"sort"
)

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

	if sortedS == sortedT {
		return true
	}
	return false
}
