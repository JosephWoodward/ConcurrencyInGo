package main

import (
	"fmt"
	"net/http"
	"time"
)

type Result struct {
	response *http.Response
	err      error
	url      string
}

func main() {
	// https://stackoverflow.com/questions/30526946/time-http-response-in-go

	urls := []string{"https://stackoverflow.com", "https://google.co.uk", "https://bbc.co.uk"}

	done := make(chan struct{})
	resCount := 0

	for result := range checkStatus(done, urls...) {
		resCount++
		if result.err != nil {
			fmt.Printf("Error: %v", result.err)
			continue
		}
		fmt.Printf("Response: %v = %v\n", result.url, result.response.Status)
		if resCount == len(urls) {
			done <- struct{}{}
			break
		}
	}

	// done <- struct{}{}
	// Can you use context instead of done? Can you share it between channels

	fmt.Printf("Finished!")
}

func checkStatus(done chan struct{}, urls ...string) <-chan Result {
	results := make(chan Result)

	for _, v := range urls {
		go func(done <-chan struct{}, url string) {
			// Heavy work is already done here, it should be moved into case after done.
			time.Sleep(2 * time.Second)
			res, err := http.Get(url)
			r := Result{response: res, err: err, url: url}

			select {
			case <-done:
				return
			case results <- r:
			}
		}(done, v)
	}

	return results
}
