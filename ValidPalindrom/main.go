package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "Was it a car or a cat I saw?"
	fmt.Println(isValidPalindrom(s))
}

func isValidPalindrom(s string) bool {

	l := 0
	r := len(s) - 1
	for l <= r {
		for l < r && !isAlphaNum(rune(s[l])) {
			l++
		}
		for r > l && !isAlphaNum(rune(s[r])) {
			r--
		}
		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
			return false
		}
		l++
		r--
	}
	return true
}

func isAlphaNum(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}
