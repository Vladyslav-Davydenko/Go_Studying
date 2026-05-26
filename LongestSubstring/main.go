package main

import "fmt"

func main() {
	s := "zxyzxyz"
	fmt.Println(LongestSubstring(s))
}

func LongestSubstring(s string) int {
	mp := make(map[byte]int)
	l, res := 0, 0

	for r := 0; r < len(s); r++ {
		if idx, found := mp[s[r]]; found {
			l = max(idx+1, l)
		}
		mp[s[r]] = r
		if r-l+1 > res {
			res = r - l + 1
		}
	}
	return res
}
