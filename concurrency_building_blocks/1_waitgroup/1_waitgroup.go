package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// WaitGroups are for fanning in
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done() // Decrement wg counter by one
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2)
	}()

	wg.Wait() // Blocks until all goroutines have completed and the `wg` is zero.

	fmt.Println("All goroutines complete.")
}
