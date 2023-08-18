package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(isPalindrome2("A man, a plan, a canal: Panama"))
}

func isPalindrome0(s string) bool {
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

func isPalindrome1(s string) bool {
	i, j := 0, len(s)-1
	arr := []rune(s)

	for i < j {
		left := unicode.ToLower(arr[i])
		right := unicode.ToLower(arr[j])

		if !isLetterOrDigit(left) {
			i++
			continue
		}

		if !isLetterOrDigit(right) {
			j--
			continue
		}

		if left != right {
			return false
		}

		i++
		j--
	}

	return true
}

func isLetterOrDigit(s rune) bool {
	return unicode.IsLetter(s) || unicode.IsDigit(s)
}

func isPalindrome2(s string) bool {
	start, end := 0, len(s)-1

	for start < end {
		for start < end && !unicode.IsLetter(rune(s[start])) && (s[start] < '0' || s[start] > '9') {
			start++
		}
		for start < end && !unicode.IsLetter(rune(s[end])) && (s[end] < '0' || s[end] > '9') {
			end--
		}
		if start < end && unicode.ToUpper(rune(s[start])) != unicode.ToUpper(rune(s[end])) {
			return false
		}
		start++
		end--
	}

	return true
}
