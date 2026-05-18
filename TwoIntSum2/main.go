package main

import "fmt"

func main() {
	nums := []int{1,2,3,4,5,6}

	fmt.Println(TwoIntSum(nums, 3))
}


func TwoIntSum(nums []int, target int) [2]int {
	l, r := 0, len(nums) - 1

	for l <= r {
		sum := nums[l] + nums[r]
		if sum == target {
			return [2]int{nums[l], nums[r]}
		} else if sum < target {
			l++
		} else {
			r--
		}
	}
	return [2]int{-1, -1}
}