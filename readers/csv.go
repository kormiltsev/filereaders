package readers

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// fdir := "./data/*.json"

func FileLister(fdir string) []string {
	mifile, err := filepath.Glob(CSVstats.Dir)
	if err != nil {
		CSVstats.Err = append(CSVstats.Err, fmt.Sprintf("ERR: no csv files in directory %s, err = %s", fdir, err))
	}
	log.Println(mifile)
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

func (c *InterfaceMiwatchSleepCSV) Read() {
	c.Files = FileLister(c.Directory)
	for _, fileadr := range c.Files {
		data, err := FileReaderCSV(fileadr, c.CSVdevider)
		if err != nil {
			c.Err = append(c.Err, fmt.Sprintf("ERR CSV %s, FILE=%s", err, fileadr))
			continue
		}
		if len(data) == 0 {
			c.Err = append(c.Err, fmt.Sprintf("ERR CSV empty file, FILE=%s", err, fileadr))
			continue
		}
		for i, t := range data[0] {
			if t != c.ColumnNames[i] {
				c.Err = append(c.Err, fmt.Sprintf("ERR wrong column names in file %s", fileadr))
				continue
			}
		}
		for i := 1; i < len(data); i++ {
			c.Rows = append(c.Rows, data[i])
		}
	}
}
