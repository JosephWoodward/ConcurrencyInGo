package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	increment := func() { count++ }
	decrement := func() { count-- }

	var once sync.Once
	once.Do(increment)
	once.Do(decrement)

	fmt.Printf("Count: %d\n", count)
	/*
	 Prints Count is 1, this is because sync.Once only counts the number of times
	 Do is called, not how many times unique functions passed are called.
	*/
}
