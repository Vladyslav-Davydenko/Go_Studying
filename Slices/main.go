package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)
	sNotNil := []string{}
	fmt.Println("uninit_Not_Nil:", sNotNil, sNotNil == nil, len(sNotNil) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	fmt.Println("t == t2:", slices.Equal(t, t2))

	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	base := []int{1, 2, 3}
	s1 := base[:2]
	s2 := base[:2]
	fmt.Println(s1[0] == s2[0]) // true -> same array

	s11 := []int{1, 2}              // literal allocates new array
	s22 := []int{1, 2}              // different array
	fmt.Println(&s11[0] == &s22[0]) // false -> different memory

	a := make([]int, 0, 2)
	b := a

	a = append(a, 1, 2, 3) // new allocation

	a[0] = 99

	fmt.Println(b) // [] — unchanged

}
