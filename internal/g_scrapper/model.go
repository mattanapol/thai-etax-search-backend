package g_scrapper

import (
	browser "github.com/EDDYCJY/fake-useragent"
)

type ScrapperResult struct {
	Title       string
	Link        string
	Description string
	Position    int
}

func getRandomUserAgent() string {
	return browser.Chrome()
}
