package readers

// Засыпание;Засыпание;Пробуждение;Пробуждение;Тип сна;Минуты сна
var MiwatchSleepCSVNames = []string{
	"Засыпание",
	"Засыпание",
	"Пробуждение",
	"Пробуждение",
	"Тип сна",
	"Минуты сна",
}
var MiWatchSleep = InterfaceCSV{
	Name:        "MiWatchSleepCSV",
	Directory:   "./data/miwatch/sleep/*.csv",
	CSVdevider:  ';',
	ColumnNames: MiwatchSleepCSVNames,
	TimeZone:    "Europe/Moscow",
	Files:       nil,
	Rows:        nil,
	FileStatus:  nil,
	Err:         nil,
}

type InterfaceCSV struct {
	Name        string
	Directory   string
	CSVdevider  rune
	ColumnNames []string
	TimeZone    string
	Files       []string
	Rows        [][]string //MiwatchSleepCSV
	FileStatus  []string
	Err         []string
}

// =================================================================
