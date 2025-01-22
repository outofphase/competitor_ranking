package main

// define datastructure:
// event ||--|{ race ||--|{ result

type Result struct {
	Position string `csv:"position"`
	Name     string `csv:"name"`
	Club     string `csv:"club"`
	Class    string `csv:"class"`
	Div      string `csv:"div"`
	Time     string `csv:"time"`
	Points   string `csv:"points"`
	Pd       string `csv:"pd"`
}

type Race struct {
	raceName string
	results  []Result
}

type Event []Race
