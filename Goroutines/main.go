package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := range 3 {
		time.Sleep(time.Millisecond)
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct")

	go f("goroutine")

	go func(msg string) {
		time.Sleep(time.Millisecond * 2)
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
