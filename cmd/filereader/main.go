package main

import (
	"log"

	"github.com/kormiltsev/filereaders/app"
	"github.com/kormiltsev/filereaders/readers"
)

type fileList struct {
	MiWatchSleepCSV bool
}

var starter = fileList{
	MiWatchSleepCSV: true,
}

func main() {
	if starter.MiWatchSleepCSV {
		// miwatch sleep csv
		structureCSV := readers.MiWatchSleep
		structureCSV.Read()
		a := &app.PGwatchsleepCSV
		a.Convert(&structureCSV)
		if len(structureCSV.Err) != 0 {
			log.Println("Error : ", structureCSV.Err)
		}
		log.Println("DONE file ", structureCSV.Name)
	}
}
