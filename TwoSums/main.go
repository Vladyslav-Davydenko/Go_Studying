package main

import (
	"fmt"
)

func main() {
	nums := []int{1,2,3,4,5,6}
	result1, result2 := twoSums(nums, 5)
	fmt.Println("idxs:", result1, result2)

}

func twoSums(nums []int, target int) (int, int) {
	set := make(map[int]int)
	for idx, n := range nums {
		if prevIdx, ok := set[n]; ok {
            return prevIdx, idx
        }
		set[target-n] = idx
	}
	return -1, -1
}