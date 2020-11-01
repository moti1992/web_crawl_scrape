package main

import (
	"fmt"
	"log"
	"os"

	"web_crawler/logic"

	"github.com/spf13/cast"
)

var allowedDomains, websiteURL, file string
var maxDepth int
var recrawl bool

func init() {
	if len(os.Args) < 5 {
		log.Println("Please refer the below sample command")
		log.Println("./{bin} {websiteURL} {allowedDomains in csv} {maxDepth to crawl} {file to save} <optional>{recrawl if not already present in the processed file}")
		return
	}

	websiteURL = os.Args[1]
	allowedDomains = os.Args[2]
	maxDepth = cast.ToInt(os.Args[3])
	file = os.Args[4]
	recrawl = false

	if len(os.Args) > 5 {
		recrawl = cast.ToBool(os.Args[5])
	}
}

func main() {
	filePath := "processed/" + file
	if !logic.FileExists(filePath) || recrawl {
		allWords := logic.CrawlAndScrape(websiteURL, allowedDomains, maxDepth)
		logic.WriteToFile(allWords, filePath)
	} else {
		fmt.Println("website crawled data result exists already")
	}

	wordsCount, err := logic.WordCount(filePath)
	if err != nil {
		log.Println("Error while processing the request. Please try again later.")
		return
	}
	logic.PrintTop10WordsAndItsCounts(wordsCount)
}
