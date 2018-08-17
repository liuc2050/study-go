package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	fmt.Println("Commencing countdown. Press return to abort.")
	go countEcho(10 * time.Second)
	select {
	case <-time.After(10 * time.Second):
		//Do nothing.
	case <-abort:
		fmt.Println("Lauch aborted!")
		return
	}
	lauch()
}

func countEcho(delay time.Duration) {
	tick := time.Tick(1 * time.Second)
	for i := delay / (1 * time.Second); i >= 1; i-- {
		fmt.Printf("%-10d\r", i)
		<-tick
	}
}

func lauch() {
	fmt.Printf("\nLauch!\n")
}
