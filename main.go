package main

import (
	"github.com/gocolly/colly"
)

type Mmo struct {
	mmo string
	url string
}

func main() {
	var mmoSaver []Mmo
	c := colly.NewCollector()

	c.OnHTML("h2", func(e *colly.HTMLElement) {
		if e.ChildText("a") != "" {
			var x Mmo
			x.mmo = e.ChildText("a")
			x.url = e.ChildAttrs("a", "href")[0]
			mmoSaver = append(mmoSaver, x)
		}
	})

	c.Visit("http://mmorpgbr.com.br")
}
