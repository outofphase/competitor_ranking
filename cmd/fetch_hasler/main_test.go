package main

import (
	"fmt"
	"os"
	"testing"
)

func TestParsePage(t *testing.T) {
	const page = "./tonbridge.html"
	var (
		result  Result
		paddler string
	)

	body, err := os.Open(page)
	if err != nil {
		t.Fatal(err)
	}
	defer body.Close()

	event, err := parsePage(body)
	if err != nil {
		t.Errorf("Failed to parse page: %v", err)
		panic(err)
	}

	got := len(event)
	want := 18
	if got != want {
		t.Errorf("Number of races incorrect: got %d, want %d", got, want)
	}

	for _, r := range event {
		if r.raceName == "div7" {
			got = len(r.results)
			want = 5
			if got != want {
				t.Errorf("Number of results incorrect: got %d, want %d", got, want)
			}
			for _, rs := range r.results {
				if rs.Position == "3" {
					result = rs
					paddler = rs.Name
				}
			}
		}
	}
	if none := (Result{"", "", "", "", "", "", "", ""}); result == none && paddler == "" {
		t.Errorf("div7 race not found")
		panic(fmt.Errorf("check file %s", page))
	}

	wantResult := Result{"3", "David Morgan", "ADS", "VMK", "7", "48", "", ""}
	if result != wantResult {
		t.Errorf("Result details incorrect: got %v, want %v", result, wantResult)
	}

	wantPaddler := "David Morgan"
	if paddler != wantPaddler {
		t.Errorf("Paddler name incorrect: got %s, want %s", paddler, wantPaddler)
	}
}
