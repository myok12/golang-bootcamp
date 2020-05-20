package main

import (
	"monitor/crawler"
	"monitor/request"
	"monitor/website"
	"time"
)

func main() {
	ch := make(chan request.Crawl, 1)
	go website.StartSite(ch)
	go crawler.CrawlServer(ch)
	time.Sleep(time.Hour*10000)
}

