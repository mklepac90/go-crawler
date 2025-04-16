package main

import (
	"fmt"
	"sort"
	"os"
	"encoding/csv"
	"strconv"
	"time"
)

type Page struct {
	URL   string
	Count int
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Printf(`
=============================
  REPORT for %s
=============================
`, baseURL)

	sortedPages := sortPages(pages)
	
	now := time.Now()
	fileName := fmt.Sprintf("report_%v.csv", now.Unix())
	file, err := os.Create(fileName)
	defer file.Close()
        if err != nil {
                fmt.Printf("failed creating file: %v", err)
        }

        writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"url", "count"})
	for _, page := range sortedPages {
		url := page.URL
		count := page.Count
		fmt.Printf("Found %d internal links to %s\n", count, url)
		writer.Write([]string{url, strconv.Itoa(count)})
	}
}

func sortPages(pages map[string]int) []Page {
	pagesSlice := []Page{}

	for url, count := range pages {
		pagesSlice = append(pagesSlice, Page{URL: url, Count: count})
	}

	sort.Slice(pagesSlice, func(i, j int) bool {
		if pagesSlice[i].Count == pagesSlice[j].Count {
			return pagesSlice[i].URL < pagesSlice[j].URL
		}

		return pagesSlice[i].Count > pagesSlice[j].Count
	})

	return pagesSlice
}
