package app

import (
	"log"
	"strconv"
	"time"

	"github.com/kormiltsev/filereaders/readers"
)

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

type MiwatchSleepCSVs struct {
	Name string
	Rows []MiwatchSleepCSV
}

var PGwatchsleepCSV = MiwatchSleepCSVs{
	Name: "MiWatchSleepCSV",
	Rows: make([]MiwatchSleepCSV, 0),
}

func (rows *MiwatchSleepCSVs) Convert(d *readers.InterfaceCSV) {
	var err error
	var row MiwatchSleepCSV
	loc, _ := time.LoadLocation(d.TimeZone)
	for _, line := range d.Rows {
		row.DateToday = time.Now()
		row.StartPeriodInt, err = strconv.Atoi(line[0])
		if err != nil {
			log.Println("Wrong int date type in column[0] ", d.Name)
		}
		layout := "02/01/2006 15:04:05" // 18/10/2020 03:15:00
		//time.ParseInLocation("2/1/2006  3:04 PM ", "31/7/2015  1:25 AM ", time.Local)

		row.StartPeriod, err = time.ParseInLocation(layout, line[1], loc)
		if err != nil {
			log.Println("Wrong date type in column[1]", d.Name)
		}
		row.EndPeriodInt, err = strconv.Atoi(line[2])
		if err != nil {
			log.Println("Wrong int date type in column[2] ", d.Name)
		}
		//layout := "02/01/2006 15:04:05" // 18/10/2020 03:15:00
		row.StartPeriod, err = time.ParseInLocation(layout, line[1], loc)
		if err != nil {
			log.Println("Wrong date type in column[3]", d.Name)
		}
		row.Dreams = line[4]
		row.PeriodDuration, err = strconv.Atoi(line[5])
		if err != nil {
			log.Println("Wrong int minutes period type in column[5] ", d.Name)
		}
		rows.Rows = append(rows.Rows, row)
	}
}

// Rows        [][]string //MiwatchSleepCSV
// FileStatus  []string
// Err         []string

// "Засыпание",
// "Засыпание",
// "Пробуждение",
// "Пробуждение",
// "Тип сна",
// "Минуты сна",
