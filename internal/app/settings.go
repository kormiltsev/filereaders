package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/kormiltsev/filereaders/internal/readers"
	"github.com/kormiltsev/filereaders/internal/storage"
)

type Settingos struct {
	PG           *storage.PGconfig
	MiWatchSleep *readers.CSVset
}

var Settings = Settingos{
	PG:           &storage.Stcon,
	MiWatchSleep: &Miwatchsleepset,
}

// func GetSettings() error {
// }
func UploadSettingsJson(s *Settingos) error {
	var err error
	file := "./settings.json"
	settingsFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
		// create new if error
		f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println("not open file")
			return err
		}
		if err := f.Close(); err != nil {
			log.Println("cant close file")
			return err
		}
		// catalog := Ws
		// return catalog, ok
	}
	defer settingsFile.Close()
	jsonParser := json.NewDecoder(settingsFile)
	if err := jsonParser.Decode(s); err != nil {
		log.Println("cant decode json")
		return err
	}
	log.Println("Catalog uploaded from: ", file)
	return nil
}

func CreateSettingsJson(s *Settingos) error {
	var err error
	file := "./settings.json"
	// d := struct {
	// 	A int
	// }{A: 2}
	rawDataOut, err := json.MarshalIndent(&Settings, "", "  ")
	log.Println(string(rawDataOut))
	if err != nil {
		log.Println("cant marshal in json")
		return err
	}
	err = ioutil.WriteFile(file, rawDataOut, 0644)
	if err != nil {
		log.Println("cant write file")
		return err
	}
	return nil
}
