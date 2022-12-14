package app

import (
	"log"
	"strconv"
	"time"

	"github.com/kormiltsev/filereaders/internal/readers"
	"github.com/kormiltsev/filereaders/internal/storage"
)

// main struct =================================================
type MiWatchSleepStruct struct {
	ToDo   bool
	Reader *readers.InterfaceCSV
	PG     []storage.MiwatchSleepRow // большой массив, а нужен ли он?? указатели не катят, работаю в одной переменной
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
	row := storage.NewMiwatchSleepRow()
	//loc, _ := time.LoadLocation(w.Reader.Settings.TimeZone)
	for _, line := range w.Reader.Rows {
		row.DateToday = time.Now()
		//row.StartPeriodInt, err = strconv.Atoi(line[0]) //int
		// check for format milliseconds
		if len(line[0]) != 13 {
			log.Println("Wrong int date. Expect in milliseconds in column[0] ", w.Reader.Settings.Name)
			return
		}
		row.StartPeriodInt, err = strconv.ParseInt(line[0], 10, 64)
		if err != nil {
			log.Println("Wrong int64 date type in column[0] ", w.Reader.Settings.Name)
		}
		layout := "02/01/2006 15:04:05" // 18/10/2020 03:15:00
		//time.ParseInLocation("2/1/2006  3:04 PM ", "31/7/2015  1:25 AM ", time.Local)
		tz, err := ReturnTimezone(row.StartPeriodInt / 1000)
		if err != nil {
			log.Println("cant find timezone history in storage ", w.Reader.Settings.Name)
			return
		}
		loc, _ := time.LoadLocation(tz)
		row.StartPeriod, err = time.ParseInLocation(layout, line[1], loc)
		if err != nil {
			log.Println("Wrong date type in column[1]", w.Reader.Settings.Name)
		}
		//row.EndPeriodInt, err = strconv.Atoi(line[2]) //int
		// check for format milliseconds
		if len(line[2]) != 13 {
			log.Println("Wrong int date. Expect in milliseconds in column[2] ", w.Reader.Settings.Name)
			return
		}
		row.EndPeriodInt, err = strconv.ParseInt(line[2], 10, 64)
		if err != nil {
			log.Println("Wrong int64 date type in column[2] ", w.Reader.Settings.Name)
		}
		//layout := "02/01/2006 15:04:05" // 18/10/2020 03:15:00
		row.EndPeriod, err = time.ParseInLocation(layout, line[3], loc)
		if err != nil {
			log.Println("Wrong date type in column[3]", w.Reader.Settings.Name)
		}
		row.Dreams = line[4]
		//row.PeriodDuration, err = strconv.Atoi(line[5]) //int
		row.PeriodDuration, err = strconv.ParseInt(line[5], 10, 64)
		if err != nil {
			log.Println("Wrong int64 minutes period type in column[5] ", w.Reader.Settings.Name)
		}
		w.PG = append(w.PG, *row)
		// go PG
		row.AddIfNotExist()

	}

}
