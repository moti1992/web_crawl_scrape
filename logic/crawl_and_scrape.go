package logic

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"golang.org/x/net/html"
)

// CrawlAndScrape ... crawl the web links till given depth, scrape the text and return the array of strings
func CrawlAndScrape(websiteURL, allowedDomains string, maxDepth int) (allWords []string) {
	log.Println("=============start=================")
	start := time.Now()

	// create the collector obj with allowedDomains, maxDepth and async to use goroutines
	// stores urls in in-memory
	c := colly.NewCollector(
		colly.AllowedDomains(allowedDomains),
		colly.MaxDepth(maxDepth),
		colly.Async(true), //recursively calling Collector.Visit from callbacks produces constantly growing stack, so async is true
	)

	// DisableKeepAlives - true if we do not want single TCP connection to remain open for multiple HTTP calls
	// c.WithTransport(&http.Transport{
	// 	DisableKeepAlives: true,
	// })

	// callback for links on scraped pages
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// extract the linked URL from the anchor tag
		link := e.Attr("href")
		// visit the link
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// rules for goroutines concurrency, allowed domains and delay between different domains matach
	c.Limit(&colly.LimitRule{
		DomainGlob: allowedDomains,
		Delay:      1 * time.Second,
		// Parallelism: 4,
	})

	// before visiting the page
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})

	// on each error occurs while crawling/scraping the page
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	// on each page scrape, do the following steps on the same
	c.OnResponse(func(r *colly.Response) {
		var err error
		log.Println("Visited", r.Request.URL)
		tokenizer := html.NewTokenizer(strings.NewReader(string(r.Body)))

		for {
			tt := tokenizer.Next()
			t := tokenizer.Token()

			err = tokenizer.Err()
			if err == io.EOF {
				break
			}

			switch tt {
			case html.ErrorToken:
				log.Fatal(err)
			case html.StartTagToken:
				tn, _ := tokenizer.TagName()
				if string(tn) == "script" {
					break
				}
			case html.TextToken: // process the html text and remove js, css, html code from the text
				p := strings.NewReader(strings.TrimSpace(t.Data))
				doc, _ := goquery.NewDocumentFromReader(p)

				doc.Find("script").Each(func(i int, el *goquery.Selection) {
					el.Remove()
				})

				// some filters by prefix & empty checks
				d := strings.TrimSpace(doc.Text())
				if d == "" {
					break
				}

				if filteredByPrefix(d) || filteredBySubStr(d) {
					break
				}

				// after the pre processing of data, push to allwords slice
				// log.Println(d)
				allWords = append(allWords, d)
			}
		}
	})

	c.Visit(websiteURL)
	c.Wait()

	log.Println("Time taken for crawl and scrape the text::", time.Since(start))
	log.Println("=============end=================")
	return
}

// WordCount ... count the words map with word and its count for a given file
func WordCount(filename string) (map[string]int, error) {
	var wordRegExp = regexp.MustCompile(`\pL+('\pL+)*`)
	var wordCounts = make(map[string]int)

	f, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return wordCounts, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Println(err)
			return wordCounts, err
		}
		line := strings.ToLower(scanner.Text())
		words := wordRegExp.FindAllString(line, -1)
		for _, word := range words {
			wordCounts[word]++
		}
	}
	return wordCounts, err
}
