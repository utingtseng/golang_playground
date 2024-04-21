package main

import (
	"fmt"
	"time"
)

// <-chan makes this read only to done channel
func doWork(done <-chan bool) {
	for {
		select {
		// allows parent go routine to end this via the done channel
		case <-done:
			return
		default:
			fmt.Println("DOING WORK")
		}
	}
}

// simulate long lived application
func main() {
	done := make(chan bool)
	go doWork(done)
	time.Sleep(time.Second * 10)

	// sends data to done channel and make doWork stop
	close(done)
}
