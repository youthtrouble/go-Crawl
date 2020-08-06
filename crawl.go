package main

import (
	"fmt"
	"net/http"
	"strings"
	
	"golang.org/x/net/html"
)



func getBody(url string) {
	resp, err := http.Get(url) 
	if err != nil {
		fmt.Printf("URL %s is invalid", url)
	}
	defer resp.Body.Close()
	anchors := []string{}
	page := html.NewTokenizer(resp.Body)
	for {
		tokenType := page.Next()
		if tokenType == html.ErrorToken {
			for _, anchor := range(anchors) {
				fmt.Println(anchor)
			}
		}
		token := page.Token()
		if tokenType == html.StartTagToken && token.DataAtom.String() == "a" {
			for _, KeyVal := range token.Attr {
				if KeyVal.Key == "href" {
					if strings.Contains(KeyVal.Val, "/jobs/") {
						fmt.Println("https://ng.indeed.com" + KeyVal.Val)
					}
				}
			}
		}
	}
}



func main() {
	getBody("https://ng.indeed.com/jobs-in-Lagos")
}