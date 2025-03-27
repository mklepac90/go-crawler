package main

import (
        "fmt"
        "net/http"
        "strings"
	"io"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	
	if err != nil {
		return "", fmt.Errorf("got network error: %v\n", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return "", fmt.Errorf("got HTTP error: %s\n", res.Status)
	}

	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("got non-HTML response: %s\n", contentType)
	}

	htmlBodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("couldn't read response body: %v\n", err)
	}

	htmlBody := string(htmlBodyBytes)

	return htmlBody, nil
}
