package main

import (
	"fmt"
)

func main () {
 nums := []int{1,2,3,4,5,6,5,5,6,3}

 s := sliceToSet(nums)
 for k, _ := range s {
	fmt.Println(k)
 }

}

func sliceToSet(nums []int) map[int]struct{} {
	set := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		set[v] = struct{}{}
	}
	return set
}