package main

// define event datastructure:
// event ||--|{ race ||--|{ result

type Result struct {
	Position string
	Name     string
	Club     string
	Class    string
	Div      string
	Time     string
	Points   string
	Pd       string
}

type Race struct {
	raceName string
	results  []Result
}

type Event []Race

type DenormalisedResult struct {
	RaceName string `csv:"race"`
	Position string `csv:"position"`
	Name     string `csv:"name"`
	Club     string `csv:"club"`
	Class    string `csv:"class"`
	Div      string `csv:"div"`
	Time     string `csv:"time"`
	Points   string `csv:"points"`
	Pd       string `csv:"pd"`
}
