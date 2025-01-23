package main

import (
	"os"

	"github.com/gocarina/gocsv"
)

func denormaliseEvent(e Event) []DenormalisedResult {
	var results []DenormalisedResult
	var dr DenormalisedResult

	for _, r := range e[0].results {
		dr.RaceName = e[0].raceName
		dr.Position = r.Position
		dr.Name = r.Name
		dr.Club = r.Club
		dr.Class = r.Class
		dr.Div = r.Div
		dr.Time = r.Time
		dr.Points = r.Points
		dr.Pd = r.Pd

		results = append(results, dr)
	}
	return results
}

func WriteData(data []DenormalisedResult) {
	eventFile, err := os.OpenFile("event.csv", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		panic(err)
	}
	defer eventFile.Close()

	err = gocsv.MarshalFile(data, eventFile)
	if err != nil {
		panic(err)
	}
}
