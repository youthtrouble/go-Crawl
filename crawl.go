package main

import (
	"fmt"
	"flag"
	"net/http"
	"os"
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
					fmt.Println(KeyVal.Val)
				}
			}
		}
	}
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {   
		fmt.Println("Please input a URL")  // if a URL wasn't provided as an argument
		os.Exit(1)                                // show a message and exit.
	  }
	  getBody(args[0])    
}