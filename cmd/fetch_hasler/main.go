package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"

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

	var e Event

	// parse web page

	eventSelection.Find(".tab-pane").Each(func(i int, raceSelection *goquery.Selection) { // for each Race
		var r Race
		r.raceName, _ = raceSelection.Attr("id")
		raceSelection.Find("tr[data-result-id]").Each(func(i int, resultSelection *goquery.Selection) { // for each Result
			var rs Result
			resultSelection.Find("td").Each(func(i int, cell *goquery.Selection) {
				cellText := cell.Text()
				switch i {
				case 0:
					rs.Position = cellText
				case 1:
					rs.Name = cellText
				case 2:
					rs.Club = cellText
				case 3:
					rs.Class = cellText
				case 4:
					rs.Div = cellText
				case 6:
					rs.Time = cellText
				case 7:
					rs.Points = cellText
				case 8:
					rs.Pd = cellText
				}
			})
			r.results = append(r.results, rs)
		})
		e = append(e, r)
	})

	// sort races

	sort.Slice(e, func(i, j int) bool {
		if len(e[i].raceName) < len(e[j].raceName) {
			return true
		}
		return e[i].raceName < e[j].raceName
	})

	// print results

	fmt.Println()
	for i, r := range e {
		fmt.Printf("race %d %s had %d paddlers\n %v\n\n", i+1, r.raceName, len(r.results), e[i])
	}

	WriteData(denormaliseEvent(e))
}
