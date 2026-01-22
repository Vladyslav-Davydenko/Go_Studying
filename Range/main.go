package main

import "fmt"

func main() {
	nums := []int{1,2,3,4,5}
	sum := 0

	for _, n := range nums {
		sum += n
	}
	fmt.Println("sum:", sum)

	for i, _ := range nums{
		fmt.Println("index:", i)
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range kvs{
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "go"{
		fmt.Println(i, c)
	}
}