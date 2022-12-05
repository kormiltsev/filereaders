package app

import (
	"log"
	"strconv"
	"time"

	"github.com/kormiltsev/filereaders/internal/readers"
)

// main struct =================================================
type MiWatchSleepStruct struct {
	ToDo   bool
	Reader *readers.InterfaceCSV
	PG     []MiwatchSleepRow
}

var MiWatchSleep = MiWatchSleepStruct{
	ToDo:   true,
	Reader: &MiWatchSleepCSV,
	PG:     nil,
}

// settings ====================================================
// using csv reader
// readers ======================================================
var MiwatchSleepCSVNames = []string{
	"Засыпание",
	"Засыпание",
	"Пробуждение",
	"Пробуждение",
	"Тип сна",
	"Минуты сна",
}
var Miwatchsleepset = readers.CSVset{
	Name:        "MiWatchSleepCSV",
	Directory:   "./data/miwatch/sleep/*.csv",
	CSVdevider:  ';',
	ColumnNames: MiwatchSleepCSVNames,
	TimeZone:    "Europe/Moscow",
}
var MiWatchSleepCSV = readers.InterfaceCSV{
	Settings:   &Miwatchsleepset,
	Files:      nil,
	Rows:       nil,
	FileStatus: nil,
	Err:        nil,
}

// pg =============================================================
type MiwatchSleepRow struct {
	ID             int
	DateToday      time.Time
	StartPeriodInt int
	StartPeriod    time.Time
	EndPeriodInt   int
	EndPeriod      time.Time
	Dreams         string
	PeriodDuration int
}

// type MiwatchSleepCSVs struct {
// 	Name string
// 	Rows []MiwatchSleep
// }

// var PGwatchsleepCSV = MiwatchSleepCSVs{
// 	Name: "MiwatchSleep",
// 	Rows: make([]MiwatchSleep, 0),
// }

// =======================================

func (w *MiWatchSleepStruct) Do() {
	readers.Read(w.Reader)
	log.Println("readers complete")
	var err error
	var row MiwatchSleepRow
	loc, _ := time.LoadLocation(w.Reader.Settings.TimeZone)
	for _, line := range w.Reader.Rows {
		row.DateToday = time.Now()
		row.StartPeriodInt, err = strconv.Atoi(line[0])
		if err != nil {
			log.Println("Wrong int date type in column[0] ", w.Reader.Settings.Name)
		}
		layout := "02/01/2006 15:04:05" // 18/10/2020 03:15:00
		//time.ParseInLocation("2/1/2006  3:04 PM ", "31/7/2015  1:25 AM ", time.Local)

		row.StartPeriod, err = time.ParseInLocation(layout, line[1], loc)
		if err != nil {
			log.Println("Wrong date type in column[1]", w.Reader.Settings.Name)
		}
		row.EndPeriodInt, err = strconv.Atoi(line[2])
		if err != nil {
			log.Println("Wrong int date type in column[2] ", w.Reader.Settings.Name)
		}
		//layout := "02/01/2006 15:04:05" // 18/10/2020 03:15:00
		row.StartPeriod, err = time.ParseInLocation(layout, line[1], loc)
		if err != nil {
			log.Println("Wrong date type in column[3]", w.Reader.Settings.Name)
		}
		row.Dreams = line[4]
		row.PeriodDuration, err = strconv.Atoi(line[5])
		if err != nil {
			log.Println("Wrong int minutes period type in column[5] ", w.Reader.Settings.Name)
		}
		w.PG = append(w.PG, row)
		// go PG
		row.AddIfNotExist()

	}
}
