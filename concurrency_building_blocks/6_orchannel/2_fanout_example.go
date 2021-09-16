package main

import (
	"fmt"
	"net/http"
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

	for result := range checkStatus(done, urls...) {
		if result.err != nil {
			fmt.Printf("Error: %v", result.err)
			continue
		}
		fmt.Printf("Response: %v = %v\n", result.url, result.response.Status)
	}

	<-done

	fmt.Printf("Finished!")
}

func checkStatus(done chan struct{}, urls ...string) <-chan Result {
	results := make(chan Result)

	for _, v := range urls {
		go func(done <-chan struct{}, url string) {
			res, err := http.Get(url)
			r := Result{response: res, err: err, url: url}
			select {
			case <-done:
				return
			case results <- r:
			}
			return
		}(done, v)
	}

	return results
}
