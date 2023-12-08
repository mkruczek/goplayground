package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("scrapethissite.com", "www.scrapethissite.com"),
	)

	c.OnHTML(".col-md-4.col-sm-6.col-xs-12", func(e *colly.HTMLElement) {
		title := e.ChildText("h3")
		description := e.ChildText("p")

		fmt.Printf("Title: %s\nDescription: %s\n\n", title, description)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c.Visit("https://www.scrapethissite.com/pages/")
	if err != nil {
		log.Println("Error visiting page:", err)
	}

}
