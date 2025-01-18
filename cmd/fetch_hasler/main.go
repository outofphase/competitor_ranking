package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Make HTTP request
	response, err := http.Get("https://entries.canoemarathon.org.uk/results/races/2024/tonbridge-marathon-2024")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Fatalf("status code error: %d %s\n", response.StatusCode, response.Status)
	}

	// Create a goquery document from the response body
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// ... (rest of the code)
	races := document.Find("div#races-tabs-content")
	if races.Size() == 0 {
		log.Fatal("Races tab not found")
	}

	type raceResult struct {
		position string
		name     string
		club     string
		class    string
		div      string
		time     string
		points   string
		pd       string
	}
	var results [][]raceResult

	races.Find(".tab-pane").Each(func(i int, race *goquery.Selection) { // for each Race
		// log.Println(division)
		var aRace struct {
			race    string
			results []raceResult
		}
		division, _ := race.Attr("id")
		aRace.race = division
		race.Find("tr[data-result-id]").Each(func(j int, result *goquery.Selection) { // for each Result
			// log.Print("found result")
			var rr raceResult
			result.Find("td").Each(func(k int, cell *goquery.Selection) {
				// Extract data from each cell
				cellText := cell.Text()
				// ... (process and store the cell data)
				// log.Println(cellText)
				switch k {
				case 0:
					rr.position = cellText
				case 1:
					rr.name = cellText
				case 2:
					rr.club = cellText
				case 3:
					rr.class = cellText
				case 4:
					rr.div = cellText
				case 6:
					rr.time = cellText
				case 7:
					rr.points = cellText
				case 8:
					rr.pd = cellText
				}
			})
			aRace.results = append(aRace.results, rr)
			// log.Println()
		})
		results = append(results, aRace.results)
	})
	// fmt.Println(results)
	fmt.Println()
	for i, rr := range results {
		fmt.Printf("race %d had %d paddlers %v\n\n", i, len(rr), results[i])
	}
}
