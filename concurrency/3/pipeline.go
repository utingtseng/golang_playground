package main

import (
	"fmt"
)

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
		// invoke
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		// keeps the for loop alive until the in channel is closed
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// stages for different operations
func main() {
	// input
	nums := []int{2, 3, 4, 5, 6}
	// we're using unbuffered channels, routines communicate synchronously
	// stage 1
	dataChannel := sliceToChannel(nums)
	// stage 2
	finalChannel := sq(dataChannel)
	// stage 3
	for n := range finalChannel {
		fmt.Println(n)
	}
}
