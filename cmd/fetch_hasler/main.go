package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	// access web page and locate races data

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

	// define datastructure:
	// event ||--|{ race ||--|{ result

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

	type race struct {
		raceName string
		results  []result
	}

	type event []race

	var e event

	// parse web page

	eventSelection.Find(".tab-pane").Each(func(i int, raceSelection *goquery.Selection) { // for each Race
		var r race
		r.raceName, _ = raceSelection.Attr("id")
		raceSelection.Find("tr[data-result-id]").Each(func(i int, resultSelection *goquery.Selection) { // for each Result
			var rs result
			resultSelection.Find("td").Each(func(i int, cell *goquery.Selection) {
				cellText := cell.Text()
				switch i {
				case 0:
					rs.position = cellText
				case 1:
					rs.name = cellText
				case 2:
					rs.club = cellText
				case 3:
					rs.class = cellText
				case 4:
					rs.div = cellText
				case 6:
					rs.time = cellText
				case 7:
					rs.points = cellText
				case 8:
					rs.pd = cellText
				}
			})
			r.results = append(r.results, rs)
		})
		e = append(e, r)
	})

	// print results

	fmt.Println()
	for i, r := range e {
		fmt.Printf("race %d %s had %d paddlers\n %v\n\n", i+1, r.raceName, len(r.results), e[i])
	}
}
