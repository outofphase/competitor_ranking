package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// access web page and locate races data
	response, err := http.Get("https://entries.canoemarathon.org.uk/results/races/2024/tonbridge-marathon-2024")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		response.Body.Close()
		log.Fatalf("http status code: %d %s\n", response.StatusCode, response.Status)
	}

	// parse web page
	event, err := parsePage(response.Body)
	if err != nil {
		panic(err)
	}

	// sort races
	event = sortRaces(event)

	// print results
	fmt.Println()
	for i, r := range event {
		fmt.Printf("race %d %s had %d paddlers\n %v\n\n", i+1, r.raceName, len(r.results), event[i])
	}

	// write CSV file
	WriteData(event)
}
