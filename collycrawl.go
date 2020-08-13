package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"

	"os"
	"strings"
)

type Job struct {
	Title string
	Company string
	URL string
}

func main() {

	c := colly.NewCollector(

		colly.AllowedDomains("ng.indeed.com", "indeed.com"),

		colly.CacheDir("./indeed"),
	)

	detailCollector := c.Clone()

	jobs := make([]Job, 0, 200)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		//if e.Attr("class") == "Button_1qxkboh-o_O-primary_cv02ee-o_O-md_28awn8-o_O-primaryLink_109aggg" {
		//	return
		//}

		link := e.Attr("href")
		if !strings.HasPrefix(link, "/companies") || strings.Index(link, "=signup") > -1 || strings.Index(link, "=login") > -1 {
			return
		}

		e.Request.Visit(link)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})

	c.Visit("https://ng.indeed.com/jobs-in-Lagos")
}
