package crawler

import (
	"fmt"
	"monitor/emailer"
	"monitor/request"
	"net/http"
	"time"
)

// crawl returns whether the site works (200).
func crawl(url string) (bool, error) {
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	return resp.StatusCode==200, nil
}

func CrawlServer(ch chan request.Crawl) {
	for {
		newRequest := <-ch
		go crawlForever(newRequest.Email, newRequest.Url)
	}
}

func crawlForever(email, url string) {
	fmt.Printf("Will crawl %v forever and notify %v on status changes", url, email)
	working := crawledOk(url)
	for {
		time.Sleep(time.Second)
		nowWorking := crawledOk(url)
		if working != nowWorking {
			emailer.SendEmail(email, url, !nowWorking)
		}
		working = nowWorking
	}
}

func crawledOk(url string) bool {
	working, err := crawl(url)
	if err != nil {
		return false
	}
	return working
}
