package readers

type FileTypeSwitcher interface {
	Read()
}

type InterfaceMiwatchSleepCSV struct {
	Name        string
	Directory   string
	CSVdevider  rune
	ColumnNames []string
	Files       []string
	Rows        [][]string //MiwatchSleepCSV
	FileStatus  []string
	Err         []string
}
