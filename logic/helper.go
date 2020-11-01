package logic

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

var blacklistPrefixStrings = []string{".", "@", "#", "var ", "(function()", "function ", "img.", "if(", "if (", "window.", "{\""}
var blacklistSubStrings = []string{"var ", "function ", "img.", "if(", "if (", "window."}

// TOP10 ... denotes 10
const TOP10 = 10

// FileExists ... check if file already exists
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Filter ... more filter if required...
func Filter(allWords []string) (filteredWords []string) {
	return
}

func filteredByPrefix(s string) bool {
	for _, val := range blacklistPrefixStrings {
		if strings.HasPrefix(s, val) {
			return true
		}
	}
	return false
}

func filteredBySubStr(s string) bool {
	for _, val := range blacklistSubStrings {
		if strings.Contains(s, val) {
			return true
		}
	}
	return false
}

// WriteToFile ... write array of string to a file
func WriteToFile(content []string, path string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println("Failed in creating file::", err)
		return err
	}
	defer file.Close()

	datawriter := bufio.NewWriter(file)
	for _, data := range content {
		datawriter.WriteString(data + "\n")
	}
	defer datawriter.Flush()
	return nil
}

func PrintTop10WordsAndItsCounts(wordsCount map[string]int) {
	log.Println("Top 10 words and its count")
	// used to switch key and value
	reverseMap := map[int]string{}
	keys := []int{}
	for key, val := range wordsCount {
		reverseMap[val] = key
		keys = append(keys, val)
	}
	// sort.Ints(keys)
	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})

	log.Println("Word --> Count")
	log.Println("===========================")
	for i, val := range keys {
		if i > TOP10-1 {
			break
		}
		log.Println(reverseMap[val], "-->", val)
	}
}
