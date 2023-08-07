package main

//-----------------------------------
import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	re := regexp.MustCompile("[^a-zA-Z0-9]+")
	cleaned := re.ReplaceAllString(s, "")
	lowerCase := strings.ToLower(cleaned)

	if lowerCase == "" {
		return true
	}

	sliceString := []rune(lowerCase)
	for i, j := 0, len(sliceString)-1; i < j; i, j = i+1, j-1 {
		sliceString[i], sliceString[j] = sliceString[j], sliceString[i]
	}
	return (string(sliceString) == lowerCase)
}

//-----------------------------------

func main() {
	isPalindrome("A man, a plan, a canal: Panama")
}
