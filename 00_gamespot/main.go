package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

type item struct {
	PageNumber int    `json:"PageNumber"`
	Title      string `json:"title"`
	Time       string `json:"time"`
	Comment    int    `json:"comment"`
	Like       int    `json:"like"`
}

func main() {

	var allItem []item
	url := "https://www.gamespot.com/news/?page=1"

	PageNumber := 2

	for i := 1; i < PageNumber; i++ {

		c := colly.NewCollector()
		c.OnHTML("div[class='card-item__content ']", func(e *colly.HTMLElement) {

			title := e.ChildText("h4[class='card-item__title ']")
			time := e.ChildText("time[class='text-small ']")
			comment, err := strconv.Atoi(e.ChildText("div:nth-child(2) > div:nth-child(1) > div:nth-child(3) > div:nth-child(2) > span:nth-child(2)"))
			if err != nil {
				fmt.Println("Can't convert this to an int!")
			}

			like, err := strconv.Atoi(e.ChildText("div:nth-child(2) > div:nth-child(1) > div:nth-child(3) > div:nth-child(3) > span:nth-child(2)"))
			if err != nil {
				fmt.Println("Can't convert this to an int!")
			}

			scrped := item{
				PageNumber: i,
				Title:      title,
				Time:       time,
				Comment:    comment,
				Like:       like,
			}

			allItem = append(allItem, scrped)
		})

		// update url
		url = "https://www.gamespot.com/news/?page=" + strconv.Itoa(i)
		c.Visit(url)
		fmt.Println(i, PageNumber)

	}

	// Write json file
	file, err := json.Marshal(allItem)
	if err != nil {
		log.Println("Unable to create json file")
	}
	os.WriteFile("export.json", file, 0644)

	fmt.Println(allItem)

}
