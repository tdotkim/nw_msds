package scraper

import (
	"strings"

	"github.com/gocolly/colly"
)

type JSONoutput struct {
	Url   string `json:"url"`
	Word  string `json:"word"`
	Count int    `json:"count"`
}

func Scrape(inputurl string, word string) JSONoutput {

	c := colly.NewCollector()

	bundle := JSONoutput{}

	c.OnHTML("body", func(e *colly.HTMLElement) {

		bundle.Url = e.Request.URL.String()
		bundle.Word = word
		bundle.Count = strings.Count(e.Text, word)

	})

	c.Visit(inputurl)

	return bundle
}
