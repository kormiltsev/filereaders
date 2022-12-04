package readers

import "time"

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

// Засыпание;Засыпание;Пробуждение;Пробуждение;Тип сна;Минуты сна
var MiwatchSleepCSVNames = []string{
	"Засыпание",
	"Засыпание",
	"Пробуждение",
	"Пробуждение",
	"Тип сна",
	"Минуты сна",
}
var MiWatchSleep = InterfaceMiwatchSleepCSV{
	Name:        "MiWatchSleepCSV",
	Directory:   "./data/miwatch/sleep/*.csv",
	CSVdevider:  ';',
	ColumnNames: MiwatchSleepCSVNames,
}
