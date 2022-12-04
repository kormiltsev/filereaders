package main

import (
	"log"
	"time"

	"github.com/kormiltsev/filereaders/readers"
)

func main() {
	// miwatch sleep csv
	structureCSV := readers.MiWatchSleep
	structureCSV.Read()
	//PGwatchsleepCSV.Convert(structureCSV)
	log.Println(structure.Rows)
}

var PGwatchsleepCSV = MiwatchSleepCSV{}

type MiwatchSleepCSV struct {
	ID             int
	DateToday      time.Time
	StartPeriodInt int
	StartPeriod    time.Time
	EndPeriodInt   int
	EndPeriod      time.Time
	Dreams         string
	PeriodDuration int
}
