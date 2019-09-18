package main

import (
	"fmt"
	"net/http"
)

func errorTest1 () {
	checkStatus := func(done <-chan interface{}, urls ...string, ) <-chan *http.Response {
		responses := make(chan *http.Response)
		go func() {
			defer close(responses)
			for _, url := range urls {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println("[Error] ->", err)
					continue
				}
				select {
				case <-done:
					return
				case responses <- resp:
				}
			}
		}()
		return responses
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://www.baidu.com", "https://badhost"}
	for response := range checkStatus(done, urls...) {
		fmt.Printf("Response: %v\n", response.Status)
	}
}

type Result struct {
	Error error
	Response *http.Response
}

func errorTest2 () {
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {

		results := make(chan Result)
		go func() {
			defer close(results)
			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{Error: err, Response: resp}
				select {
				case <-done:
					return
				case results <- result:
				}
			}
		}()
		return results
	}

	done := make(chan interface{})
	defer close(done)

	errorCount := 0
	urls := []string{"https://www.baidu.com", "https://badhost", "https://www.wechat.com", "localhost", "192.168.0.1", "192.168.1.1", "119.23.33.220"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("[Error] -> %v\n", result.Error)
			errorCount++
			if errorCount >= 3 {
				fmt.Println("Too many errors, breaking!")
				break
			}
			continue
		}
		errorCount = 0
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}

func main () {
	errorTest2()
}