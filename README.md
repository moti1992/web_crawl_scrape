# Web crawler and scraper

This will crawl the web page with given depth, scrape those urls, then write the processed/filtered text contents to a file
Then file will be passed to do further tasks like counting of word, etc...
Here we have implemented the word count

### Technology & Tools

Project uses:

* [Golang] - Programming language
* [html] - golang.org/x/net/html - parse the html web content

Libraries:
* [Goquery] - github.com/PuerkitoBio/goquery - query the scraped data
* [Colly] - github.com/gocolly/colly - to crawl, scrape the data

### Installation

Install Golang:
Guide: https://medium.com/@patdhlk/how-to-install-go-1-9-1-on-ubuntu-16-04-ee64c073cd79

Install the dependencies:

```sh
change to project directory and run the below command
go get ./
```

### How to Build and execute

```sh
Run 
go build main.go
./main {websiteURL} {allowedDomains in csv} {maxDepth to crawl} {file to save} <optional>{recrawl if not already present in the processed file}

Example:
./main "https://www.314e.com/" "www.314e.com" 4 314e false
```





