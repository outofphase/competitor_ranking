package main

import (
	"os"

	"github.com/gocarina/gocsv"
)

func denormalise(event Event) []DenormalisedResult {
	var (
		results []DenormalisedResult
		dr      DenormalisedResult
	)
	for _, r := range event {
		dr.RaceName = r.raceName
		for _, rs := range r.results {
			dr.Position = rs.Position
			dr.Name = rs.Name
			dr.Club = rs.Club
			dr.Class = rs.Class
			dr.Div = rs.Div
			dr.Time = rs.Time
			dr.Points = rs.Points
			dr.Pd = rs.Pd

			results = append(results, dr)
		}
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
