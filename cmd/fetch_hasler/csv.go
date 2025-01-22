package main

import (
	"github.com/gocarina/gocsv"
	"os"
)

func WriteData(data []Result) {

	eventFile, err := os.OpenFile("event.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer eventFile.Close()

	err = gocsv.MarshalFile(data, eventFile) // Use this to save the CSV back to the file
	if err != nil {
		panic(err)
	}

}
