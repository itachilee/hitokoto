package crawl

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/gookit/goutil/dump"
)

func CrawlNovel() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("#main > div:nth-child(2) > div.recomclass > dl:nth-child(1) > dd.tit", func(e *colly.HTMLElement) {
		dump.P(e.Text)
		// e.Request.Visit(e.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("http://23qb.com/")
}
