package main

import (
	"fmt"
	"strconv"

	"github.com/gocolly/colly"
)

func main() {
	site := "https://www.namasha.com/Lamourlove"
	pageNumber := 20
	url := site
	lastId := "0"
	allLinks := make([]string, 0)

	for i := 1; i < pageNumber; i++ {

		c := colly.NewCollector()
		c.OnHTML("div[class='thumbnail-text pt-1 pt-md-2 pt-lg-1 mt-1']", func(e *colly.HTMLElement) {
			links := e.ChildAttr("a[class='thumbnail-title thumbnail-url flex-shrink-1 stretched-link']", "href")
			allLinks = append(allLinks, links)
		})

		c.OnHTML("div[class='thumbnail full-width mx-n2 mx-md-0 exerted-watch-time']", func(e *colly.HTMLElement) {
			dataId := e.ChildAttrs("div[class*='mx-n2 mx-md-0']", "data-id")
			lastId = dataId[59]
		})

		c.Visit(url)
		url = site + "?page=" + strconv.Itoa(i) + "&qm=&lastId=" + lastId
	}
	fmt.Println(allLinks)
}
