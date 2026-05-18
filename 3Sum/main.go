package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{-1,0,1,2,-1,-4}
	fmt.Println(ThreeSum(nums))
}

func ThreeSum(nums []int) [][]int {
	slices.Sort(nums)
	result := [][]int{}

	for i := 0; i < len(nums); i++{
		a := nums[i]
		if a > 0 {
			break
		}
		if i > 0 && a == nums[i-1] {
			continue
		}
		l, r := i + 1, len(nums) - 1
		for l < r {
			sum3 := a + nums[l] + nums[r]
			if sum3 < 0 {
				l++
			} else if sum3 > 0 {
				r--
			} else {
				result = append(result, []int{a, nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1]{
					l++
				}
			}
		}

	}
	return result
}
