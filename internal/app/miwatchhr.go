package app

import (
	"log"
	"strconv"
	"time"

	"github.com/kormiltsev/filereaders/internal/readers"
	"github.com/kormiltsev/filereaders/internal/storage"
)

// main struct =================================================
type MiWatchHrStruct struct {
	ToDo   bool
	Reader *readers.InterfaceCSV
	PG     []storage.MiwatchHrRow // большой массив, а нужен ли он?? указатели не катят, работаю в одной переменной
}

var MiWatchHr = MiWatchHrStruct{
	ToDo:   true,
	Reader: &MiWatchHrCSV,
	PG:     nil,
}

// settings ====================================================
// using csv reader
// readers ======================================================
var MiwatchHrCSVNames = []string{ //Пульс;Метка времени;Дата;Время
	"Пульс",
	"Метка времени",
	"Дата",
	"Время",
}
var Miwatchhrset = readers.CSVset{
	Name:        "MiWatchHrCSV",
	Directory:   "./data/miwatch/heartrate/*.csv",
	CSVdevider:  ';',
	ColumnNames: MiwatchHrCSVNames,
	TimeZone:    "Europe/Moscow",
}
var MiWatchHrCSV = readers.InterfaceCSV{
	Settings:   &Miwatchhrset,
	Files:      nil,
	Rows:       nil,
	FileStatus: nil,
	Err:        nil,
}

// =======================================

func (w *MiWatchHrStruct) Do() {
	readers.Read(w.Reader)
	log.Println("readers complete")
	var err error
	row := storage.NewMiwatchHrRow()
	for _, line := range w.Reader.Rows {
		row.DateToday = time.Now()
		// time in milliseconds
		row.EventTimeInt, err = strconv.Atoi(line[1])
		if err != nil {
			log.Println("Wrong int date type in column[1] ", w.Reader.Settings.Name)
		}
		row.Heartrate, err = strconv.Atoi(line[0])
		if err != nil {
			log.Println("Wrong int heartrate type in column[0] ", w.Reader.Settings.Name)
		}
		w.PG = append(w.PG, *row)
		// go PG
		row.AddIfNotExist()

	}
}
