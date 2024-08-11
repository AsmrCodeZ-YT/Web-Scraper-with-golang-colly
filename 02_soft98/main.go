package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

type Item struct {
	Title       string `json:"Titles"`
	View        string `json:"Views"`
	Writer      string `json:"Writers"`
	Categorical string `json:"Categoricals"`
	Time        string `json:"times"`
}

func main() {
	var items []Item
	pageNumber := 100 

	for i := 1; i < pageNumber; i++ {

		c := colly.NewCollector()
		// Find and print all links
		c.OnHTML("article[class='cbdd']", func(e *colly.HTMLElement) {

			title := e.ChildText("h2[class='cbddt']")
			view := e.ChildText("span[class='cbddiv']")
			writer := e.ChildText("span[class='cbddia']")
			categori := e.ChildText("span[class='cbddic']")
			time := e.ChildText("span[class='cbddid']")

			item1 := Item{
				Title:       title,
				View:        view,
				Writer:      writer,
				Categorical: categori,
				Time:        time,
			}

			items = append(items, item1)

		})
		url := "https://soft98.ir/page/" + strconv.Itoa(i)
		fmt.Println(i, url)
		c.Visit(url)
	}

	file, err := json.Marshal(items)
	if err != nil {
		log.Println("Unable to create json file")
	}
	os.WriteFile("export.json", file, 0644)

}
