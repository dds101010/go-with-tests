package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func slowDummyCrawler(url string) bool {
	time.Sleep(2 * time.Second)
	if url == "www.google.com" {
		return true
	}
	return false
}

// inside CheckWebsites function, each WebsiteCrawler functioin is blocking. hence it takes ~6 seconds
func TestWebsiteCrawler(t *testing.T) {
	urls := []string{"www.google.com", "www.facebook.com", "www.linkedin.com"}

	want := map[string]bool{
		"www.google.com":   true,
		"www.facebook.com": false,
		"www.linkedin.com": false,
	}

	got := CheckWebsites(slowDummyCrawler, urls)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

/*
[2022/08/21 03:35:37.294Z] all crawlers triggered, polling values now
[2022/08/21 03:35:37.294Z] crawling www.linkedin.com
[2022/08/21 03:35:37.294Z] crawling www.google.com
[2022/08/21 03:35:37.294Z] crawling www.facebook.com
[2022/08/21 03:35:39.295Z] got {www.facebook.com false} in channel
[2022/08/21 03:35:39.295Z] got {www.linkedin.com false} in channel
[2022/08/21 03:35:39.296Z] got {www.google.com true} in channel
*/
func TestWebsiteCrawlerOptimized(t *testing.T) {
	urls := []string{"www.google.com", "www.facebook.com", "www.linkedin.com"}

	want := map[string]bool{
		"www.google.com":   true,
		"www.facebook.com": false,
		"www.linkedin.com": false,
	}

	got := CheckWebsitesOptimized(slowDummyCrawler, urls)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
