package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 2, 3, 3, 4, 4, 4, 5, 5, 5, 5, 5, 6}
	fmt.Println(topKFreqElements(nums, 3))
}

func topKFreqElements(nums []int, k int) []int {
	count := map[int]int{}
	freq := make([][]int, len(nums)+1)
	result := []int{}

	for _, n := range nums {
		count[n]++
	}

	for num, cnt := range count {
		freq[cnt] = append(freq[cnt], num)
	}

	for i := len(freq) - 1; i > 0; i-- {
		for _, num := range freq[i] {
			result = append(result, num)
			if len(result) >= k {
				return result
			}
		}
	}
	return result
}
