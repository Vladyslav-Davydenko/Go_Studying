package main

import "fmt"

func main() {
	result := isAnagram("racecar", "carrace")
	fmt.Println(result)
}

func isAnagram(s1, s2 string) bool {
	set := make(map[rune]int)

	if len(s1) != len(s2) {
		return false
	}

	for _, value := range s1 {
		set[value]++
	}
	for _, value := range s2 {
		set[value]--
	}
	for _, value := range set {
		if value != 0 {
			return false
		}
	}

	return true
}
