package main

import (
	"errors"
	"fmt"

	"github.com/gocolly/colly/v2"
)

type Mmo struct {
	mmo string
	url string
}

func mmoRpgBr(page int) ([]Mmo, error) {
	fmt.Println(page)
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

	if len(mmoSaver) > 0 {
		return mmoSaver, nil
	}
	return mmoSaver, errors.New("No results")
}

func mmoRpg(page int) ([]Mmo, error) {
	fmt.Println(page)
	var mmoSaver []Mmo
	c := colly.NewCollector()

	c.OnHTML(".item__content__title", func(e *colly.HTMLElement) {
		if e.ChildText("a") != "" {
			var x Mmo
			x.mmo = e.ChildText("a")
			x.url = "https://www.mmorpg.com" + e.ChildAttrs("a", "href")[0]
			mmoSaver = append(mmoSaver, x)
		}
	})
	c.Visit("https://www.mmorpg.com")

	if len(mmoSaver) > 0 {
		return mmoSaver, nil
	}
	return mmoSaver, nil //errors.New("No results")
}

func mmoSteam(page int) ([]Mmo, error) {
	fmt.Println(page)
	var mmoSaver []Mmo
	c := colly.NewCollector()

	c.OnHTML("#NewReleasesRows", func(e *colly.HTMLElement) {
		var x Mmo
		for index, _ := range e.ChildTexts(".tab_item_name") {
			x.mmo = e.ChildTexts(".tab_item_name")[index]
			x.url = e.ChildAttrs("a", "href")[index]
			mmoSaver = append(mmoSaver, x)
		}
	})
	c.Visit("https://store.steampowered.com/tags/en/MMORPG/")

	if len(mmoSaver) > 0 {
		return mmoSaver, nil
	}
	return mmoSaver, nil //errors.New("No results")
}

func main() {
	mmo1, err := mmoRpgBr(1)
	if err != nil {
		fmt.Println(mmo1)
		panic(err)
	}
	mmo2, err := mmoRpg(1)
	if err != nil {
		fmt.Println(mmo2)
		panic(err)
	}
	mmo3, err := mmoSteam(1)
	if err != nil {
		fmt.Println(mmo3)
		panic(err)
	}
	fmt.Println(mmo3[0])
}

//i need to review this code.
