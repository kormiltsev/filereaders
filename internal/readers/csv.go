package readers

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

// fdir := "./data/*.json"
// main struct ==============================
type InterfaceCSV struct {
	Settings   *CSVset
	Files      []string
	Rows       [][]string
	FileStatus []string
	Err        []string
}
type CSVset struct {
	Name        string   `json:"name_csv"`
	Directory   string   `json:"directory_csv"`
	CSVdevider  rune     `json:"csv_devider"`
	ColumnNames []string `json:"column_names"`
	TimeZone    string   `json:"timezone"`
}

// ===========================================
func Read(c *InterfaceCSV) {
	var err error
	c.Files, err = FileLister(c.Settings.Directory)
	if err != nil {
		c.Err = append(c.Err, fmt.Sprintf("ERR CSV %s, DIRECTORY=%s", err, c.Settings.Directory))
		return
	}
	for _, fileadr := range c.Files {
		data, err := FileReaderCSV(fileadr, c.Settings.CSVdevider)
		if err != nil {
			c.Err = append(c.Err, fmt.Sprintf("ERR CSV %s, FILE=%s", err, fileadr))
			continue
		}
		if len(data) == 0 {
			c.Err = append(c.Err, fmt.Sprintf("ERR CSV empty file, FILE=%s", fileadr))
			continue
		}
		for i, t := range data[0] {
			if t != c.Settings.ColumnNames[i] {
				c.Err = append(c.Err, fmt.Sprintf("ERR wrong column names in file %s", fileadr))
				continue
			}
		}
		for i := 1; i < len(data); i++ {
			c.Rows = append(c.Rows, data[i])
		}
	}
}

func FileLister(fdir string) ([]string, error) {
	mifile, err := filepath.Glob(fdir)
	if err != nil {
		return nil, err
	}
	return mifile, nil
}

func FileReaderCSV(filename string, devider rune) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := csv.NewReader(f) //strings.NewReader(f))
	r.Comma = devider
	r.Comment = '#'
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}
