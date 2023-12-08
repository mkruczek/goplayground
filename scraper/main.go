package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {
	// Tworzymy nową instancję Colly
	c := colly.NewCollector(
		colly.AllowedDomains("scrapethissite.com", "www.scrapethissite.com"),
	)

	// Ustawiamy akcję dla każdego odwiedzonego elementu
	c.OnHTML(".col-md-4.col-sm-6.col-xs-12", func(e *colly.HTMLElement) {
		title := e.ChildText("h3")
		description := e.ChildText("p")

		fmt.Printf("Title: %s\nDescription: %s\n\n", title, description)
	})

	// Obsługa błędów
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Rozpoczęcie skrapowania
	err := c.Visit("https://www.scrapethissite.com/pages/")
	if err != nil {
		log.Println("Error visiting page:", err)
	}

}
