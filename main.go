package main

import (
	"fmt"
	"log"
	"os"

	"web_crawler/logic"

	"github.com/spf13/cast"
)

var allowedDomains, websiteURL string
var maxDepth int
var recrawl bool

func init() {
	if len(os.Args) < 4 {
		log.Println("Please refer the below sample command")
		log.Println("./{bin} {websiteURL} {allowedDomains in csv} {maxDepth to crawl}")
		return
	}

	websiteURL = os.Args[1]
	allowedDomains = os.Args[2]
	maxDepth = cast.ToInt(os.Args[3])
	recrawl = false

	if len(os.Args) > 4 {
		recrawl = cast.ToBool(os.Args[4])
	}
}

func main() {
	if logic.FileExists("processed/" + websiteURL) {
		fmt.Println("website file exists")
	}
	allWords := logic.CrawlAndScrape(websiteURL, allowedDomains, maxDepth)
	log.Println("Words before filtered::", allWords)
	filteredWords := logic.Filter(allWords)
	log.Println("Words after filtered::", filteredWords)
}
