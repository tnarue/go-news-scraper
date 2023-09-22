package main

import (
	"fmt"
	"gosample/objects"

	"github.com/gocolly/colly"
)

func headlineScraper(url string) []objects.Headline {
	newsList := []objects.Headline{}
	c := colly.NewCollector(
		colly.AllowedDomains(url),
	)

	c.OnHTML("section.story-wrapper", func(e *colly.HTMLElement) {
		var headline objects.Headline
		e.ForEach("a", func(_ int, e *colly.HTMLElement) {

			headline.Url = e.Attr("href")
			headline.Text = e.ChildText("h3")
			newsList = append(newsList, headline)

			fmt.Printf("Headline: %s \n Link: %s \n\n", headline.Text, headline.Url)
		})

	})
	c.Visit("https://" + url)
	return newsList
}
func main() {
	targetedUrl := "www.nytimes.com"
	newsList := headlineScraper(targetedUrl)

	fmt.Printf("Number of Headline News on front page: %v", len(newsList))
}
