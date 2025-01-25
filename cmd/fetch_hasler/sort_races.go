package main

import "sort"

func sortRaces(e Event) Event {
	sort.Slice(e, func(i, j int) bool {
		// sort doubles races after singles sortRaces
		if len(e[i].raceName) < len(e[j].raceName) {
			return true
		}
		return e[i].raceName < e[j].raceName
	})
	return e
}
