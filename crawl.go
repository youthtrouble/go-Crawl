package main

import (
	"fmt"
	"net/http"
	"strings"
	//"flag"
	//"os"
	
	"golang.org/x/net/html"
)



func getBody(url string) {
	resp, err := http.Get(url) 
	if err != nil {
		fmt.Printf("URL %s is invalid", url)
	}
	defer resp.Body.Close()

	anchors := make(map[string]bool)

	page := html.NewTokenizer(resp.Body)
	for {
		tokenType := page.Next()
		if tokenType == html.ErrorToken {
			// for _, anchor := range(anchors) {
			// 	fmt.Println(anchor)
			break
			
		}

		token := page.Token()
		if tokenType == html.StartTagToken && token.DataAtom.String() == "a" {
			for _, KeyVal := range token.Attr {
				if KeyVal.Key == "href" {
					if strings.Contains(KeyVal.Val, "/jobs/") || strings.Contains(KeyVal.Val, "/pagead/") || strings.Contains(KeyVal.Val, "/rc/") {
						url := "https://ng.indeed.com" + KeyVal.Val
						_, exists := anchors[url]
						if !exists {
							anchors[url] = true
						}
						break
					}
					//if strings.Contains(KeyVal.Val, "/jobs?q=&l=Lagos&start") {
					//	next := "https://ng.indeed.com" + KeyVal.Val
					//	getBody(next)
					//}
				}

			}
		}
	}
	for key := range anchors {
		fmt.Println(key)
	}
}



func main() {
	// flag.Parse()
	// args := flag.Args()
	// if len(args) < 1 {   
	// 	fmt.Println("Please specify start page")  // if a starting page wasn't provided as an argument
	// 	os.Exit(1)                                // show a message and exit.
	//   }
	//   getBody(args[0])
	getBody("https://ng.indeed.com/jobs-in-Lagos")
}