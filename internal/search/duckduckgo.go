package search

import (
	"fmt"
	"net/url"
	"time"

	"github.com/gocolly/colly/v2"
)

type SearchResult struct {
	Title   string
	Content string
	URL     string
}

func SearchDuckDuckGo(query string) ([]SearchResult, error) {
	results := []SearchResult{}
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
	)

	c.SetRequestTimeout(20 * time.Second)
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 1,
		Delay:       1 * time.Second,
	})

	c.OnHTML("div.result", func(e *colly.HTMLElement) {
		result := SearchResult{
			Title:   e.ChildText("h2"),
			Content: e.ChildText("div.snippet"),
			URL:     e.ChildAttr("a.result__url", "href"),
		}
		results = append(results, result)
	})

	searchURL := fmt.Sprintf("https://html.duckduckgo.com/html/?q=%s", url.QueryEscape(query))
	err := c.Visit(searchURL)
	if err != nil {
		return nil, err
	}

	return results, nil
}