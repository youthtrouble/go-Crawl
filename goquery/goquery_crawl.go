package goquery

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func summonQuery() {
	doc, err := goquery.NewDocument("https://ng.indeed.com/jobs-in-Lagos")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find()
}
