package concurrency

import (
	"fmt"
	"time"
)

// WebsiteCrawler is any function that takes a string and returns a boolean
type WebsiteCrawler func(string) bool

func CheckWebsites(wc WebsiteCrawler, urls []string) map[string]bool {
	result := map[string]bool{}

	for _, url := range urls {
		result[url] = wc(url)
	}

	return result
}

type crawlerResponse struct {
	string
	bool
}

func CheckWebsitesOptimized(wc WebsiteCrawler, urls []string) map[string]bool {
	result := map[string]bool{}
	resultChannel := make(chan crawlerResponse)

	for _, url := range urls {

		// anonymous function func(...){...}(...), this is similar to JavaScript
		// channel <- value = push value to channel
		// function execution preceeded with go runs non-blocking. known as `goroutine`
		go func(url string) {
			log(fmt.Sprintf("crawling %v", url))
			resultChannel <- crawlerResponse{url, wc(url)}
		}(url)
	}

	log("all crawlers triggered, polling values now")
	for i := 0; i < len(urls); i++ {
		// variable := <- channel = pulling the value out of channel. Note: this is blocking
		// hence, if there is no value in channel, the execution will halt here until a value is pushed
		r := <-resultChannel
		log(fmt.Sprintf("got %v in channel", r))
		result[r.string] = r.bool
	}

	return result
}

func log(str string) {
	// https://stackoverflow.com/a/20234207 -- for time formatting
	t := time.Now().Format("2006/01/02 03:04:05.000Z")
	fmt.Printf("[%v] %v\n", t, str)
}
