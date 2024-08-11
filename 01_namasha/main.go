package main

import (
	"fmt"
	"strconv"

	"github.com/gocolly/colly"
	"github.com/mavihq/persian"
)

type Item struct {
	Title   string `json:"Titles"`
	Like    int    `json:"Like"`
	Comment string `json:"Writers"`
	Time    string `json:"times"`
}

func main() {
	// var items []Item

	c := colly.NewCollector()
	// Find and print all links
	c.OnHTML("div[class='col-lg-7 col-xl-8 px-0 px-lg-3']", func(e *colly.HTMLElement) {

		title := e.ChildText("h1[class='video-title font-size-base font-size-lg-lg font-weight-bold mb-1']")
		like, _ := strconv.Atoi(persian.ToEnglishDigits(e.ChildText("span[class='video-like-count']")))
		commentC := e.ChildText("span[id='comment-count']")
		time := e.ChildText("time[class='text-dynamic-half-dark font-size-xs font-weight-light font-weight-lg-medium mt-2']")

		item1 := Item{
			Title:   title,
			Like:    like,
			Comment: commentC,
			Time:    time,
		}
		
		fmt.Println(item1)
	})
	url := "https://www.namasha.com/v/fBiBpJEB"
	c.Visit(url)
}
