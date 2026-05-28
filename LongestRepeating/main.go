package main

import "fmt"

func main() {
	s := "AAABABB"
	fmt.Println(characterReplacement(s, 1))
}

func LongestRepeatingWithReplacement(s string, k int) int {
	charSet := make(map[byte]bool)
	res := 0

	for i := 0; i < len(s); i++ {
		charSet[s[i]] = true
	}

	for c := range charSet {
		count, l := 0, 0
		for r := 0; r < len(s); r++ {
			if s[r] == c {
				count++
			}

			for (r-l+1)-count > k {
				if s[l] == c {
					count--
				}
				l++
			}

			res = max(res, r-l+1)
		}
	}
	return res
}

func characterReplacement(s string, k int) int {
	count := make(map[byte]int)
	res, l, maxf := 0, 0, 0

	for r := 0; r < len(s); r++ {
		count[s[r]]++
		if count[s[r]] > maxf {
			maxf = count[s[r]]
		}

		for (r-l+1)-maxf > k {
			count[s[l]]--
			l++
		}

		if r-l+1 > res {
			res = r - l + 1
		}
	}

	return res
}
