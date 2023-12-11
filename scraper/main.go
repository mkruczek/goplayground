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

	var counter int

	c.OnHTML("div.col-md-4.country", func(e *colly.HTMLElement) {
		title := e.ChildText("h3.country-name")
		capital := e.ChildText("span.country-capital")
		population := e.ChildText("span.country-population")
		area := e.ChildText("span.country-area")

		fmt.Printf("Title: %s\nDescription: %s // %s // %s\n\n", title, capital, population, area)
		counter++
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c.Visit("https://www.scrapethissite.com/pages/simple")
	if err != nil {
		log.Println("Error visiting page:", err)
	}

	fmt.Printf("Scraped %d pages\n", counter)
}
