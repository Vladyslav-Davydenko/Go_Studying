package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	strs := []string{"act", "pots", "tops", "cat", "stop", "hat"}
	fmt.Println(groupAnagrams(strs))
}

func groupAnagrams(strs []string) [][]string {
	set := make(map[[26]int][]string)
	for _, str := range strs {
		key := [26]int{}
		for _, r := range str {
			idx := r - 'a' // automatically converts to 97
			key[idx]++
		}
		set[key] = append(set[key], str)
	}
	return slices.Collect(maps.Values(set))
}
