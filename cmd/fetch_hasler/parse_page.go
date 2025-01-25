package main

import (
	"errors"
	"io"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func parsePage(body io.ReadCloser) (e Event, err error) {
	document, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Println("Error loading HTTP response body.")
		return nil, err
	}

	eventSelection := document.Find("div#races-tabs-content")
	if eventSelection.Size() == 0 {
		log.Println("Races tab not found.")
		return nil, errors.New("Races tab not found.")
	}
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

	return e, nil
}
