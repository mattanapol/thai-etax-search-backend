package g_scrapper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
	"io"
	"time"
)

type Service interface {
	Scrap(keyword string) ([]ScrapperResult, error)
}

type service struct {
}

func (s *service) Scrap(keyword string) ([]ScrapperResult, error) {
	client := resty.New()
	client.AddRetryCondition(func(response *resty.Response, err error) bool {
		return response.StatusCode() == 429
	},
	)
	client.SetRetryWaitTime(10 * time.Minute)
	client.SetRetryCount(10)
	resp, err := client.R().
		SetDoNotParseResponse(true).
		SetHeaders(map[string]string{
			"User-Agent": getRandomUserAgent(),
		},
		).
		SetQueryParams(map[string]string{
			"q": keyword,
		},
		).
		Get("https://www.google.com/search")
	if err != nil {
		return []ScrapperResult{}, err
	}
	if resp.IsError() {
		return []ScrapperResult{}, fmt.Errorf("error while scrapping google: %s", resp.Status())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.RawResponse.Body)
	return scrapResultFromHtml(resp.RawResponse.Body)
}

func scrapResultFromHtml(htmlResultReader io.Reader) ([]ScrapperResult, error) {
	var resultList []ScrapperResult
	doc, err := goquery.NewDocumentFromReader(htmlResultReader)
	if err != nil {
		return []ScrapperResult{}, err
	}
	c := 0
	doc.Find("div.g").Each(func(i int, result *goquery.Selection) {
		title := result.Find("h3").First().Text()
		link, _ := result.Find("a").First().Attr("href")
		snippet := result.Find(".VwiC3b").First().Text()

		resultList = append(resultList, ScrapperResult{
			Title:       title,
			Link:        link,
			Description: snippet,
			Position:    c + 1,
		},
		)
		c++
	},
	)

	return resultList, nil
}
