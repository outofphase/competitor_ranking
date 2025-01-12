package main

import (
	//	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Make HTTP request
	response, err := http.Get("https://entries.canoemarathon.org.uk/results/races/2024/tonbridge-marathon-2024/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create a goquery document from the response body
	// document, err := goquery.NewDocumentFromReader(response.Body)
	_, err = goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// ... (rest of the code)
}
