package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	//"flag"
	//"os"
	
	"golang.org/x/net/html"
)

type jobInfo struct {
	Title string
}

func getBody(url string, nextURLs map[string]bool) {

	if len(nextURLs) > 1 {
		return
	}

	resp, err := http.Get(url) 
	if err != nil {
		fmt.Printf("URL %s is invalid", url)
	}
	defer resp.Body.Close()

	//var Info string
	anchors := make(map[string]jobInfo)

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
						if exists {
							continue
						}

						//going into the links to obtain information
						jobresp, err := http.Get(url)
						if err != nil {
							log.Println("Job link unresponsive")
							continue
						}

						defer jobresp.Body.Close()

						//tokenize the page
						jobpage := html.NewTokenizer(jobresp.Body)

						for {
							jobTokentype := jobpage.Next()

							if jobTokentype == html.ErrorToken {
								break
							}

							//if we find a h1 tag, extract the data
							jobtoken := jobpage.Token()
							if jobTokentype == html.StartTagToken && token.DataAtom.String() == "h1" {
								anchors[url] = jobInfo{Title: jobtoken.Data}
							}
						}
					}
					if strings.Contains(KeyVal.Val, "/jobs?q=&l=Lagos&start") {
						next := "https://ng.indeed.com" + KeyVal.Val
						_, URLexists := nextURLs[next]
						if !URLexists {
							nextURLs[next] = true
						} else {
							continue
						}
						for key := range anchors {
							fmt.Println(key)
						}
						log.Println("Going to", next)
						getBody(next, nextURLs)
					}
				}

			}
		}
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
	getBody("https://ng.indeed.com/jobs-in-Lagos", make(map[string]bool))
}