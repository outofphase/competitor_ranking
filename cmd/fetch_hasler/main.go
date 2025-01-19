package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	response, err := http.Get("https://entries.canoemarathon.org.uk/results/races/2024/tonbridge-marathon-2024")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Fatalf("status code error: %d %s\n", response.StatusCode, response.Status)
	}

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	eventSelection := document.Find("div#races-tabs-content")
	if eventSelection.Size() == 0 {
		log.Fatal("Races tab not found")
	}

	type result struct {
		position string
		name     string
		club     string
		class    string
		div      string
		time     string
		points   string
		pd       string
	}
	var event [][]result // event ||--|{ race ||--|{ result

	eventSelection.Find(".tab-pane").Each(func(i int, raceSelection *goquery.Selection) { // for each Race
		var race struct {
			raceName string
			results  []result
		}
		divisionID, _ := raceSelection.Attr("id")
		race.raceName = divisionID
		raceSelection.Find("tr[data-result-id]").Each(func(i int, resultSelection *goquery.Selection) { // for each Result
			var rr result
			resultSelection.Find("td").Each(func(i int, cell *goquery.Selection) {
				cellText := cell.Text()
				switch i {
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
			race.results = append(race.results, rr)
		})
		event = append(event, race.results)
	})

	fmt.Println()
	for i, rr := range event {
		fmt.Printf("race %d had %d paddlers %v\n\n", i, len(rr), event[i])
	}
}
