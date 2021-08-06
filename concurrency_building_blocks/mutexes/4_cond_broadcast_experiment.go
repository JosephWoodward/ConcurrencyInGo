package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	c := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup

	fn := func(i int) {
		defer wg.Done()

		c.L.Lock()
		time.Sleep(time.Duration(i) * time.Second)
		fmt.Printf("Go %d...", i)
		c.L.Unlock()
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go fn(i)
	}

	c.Broadcast()

	wg.Wait()
	fmt.Println("Done")

}
