package main

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	mu        sync.Mutex
	items     []Item
	itemAdded sync.Cond
}

type Item struct {
	Value string
}

func newQueue() *Queue {
	q := new(Queue)
	q.itemAdded.L = &q.mu
	return q
}

func (q *Queue) Get() Item {
	q.mu.Lock()
	defer q.mu.Unlock()
	for len(q.items) == 0 {
		fmt.Println("Get() - waiting for signal...")
		q.itemAdded.Wait()
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue) Put(item Item) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, item)
	fmt.Println("Item added, signalling one gorountine (if it exists)")
	q.itemAdded.Signal() // Locks mutex then wakes goroutine
}

func main() {
	var wg sync.WaitGroup
	q := newQueue()

	wg.Add(2)
	go func() {
		fmt.Println("Launching a goroutine to get from queue")
		defer wg.Done()
		item := q.Get()
		fmt.Printf("Hello %s 1\n", item.Value)
	}()
	go func() {
		fmt.Println("Launching a goroutine to get from queue")
		defer wg.Done()
		item := q.Get()
		fmt.Printf("Hello %s 2\n", item.Value)
	}()

	// A bit of sleep so we can see things happening
	time.Sleep(2 * time.Second)
	q.Put(Item{"World"}) // Wakes goroutine 1

	time.Sleep(2 * time.Second)
	q.Put(Item{"World"}) // Wakes goroutines 2

	wg.Wait()
}
