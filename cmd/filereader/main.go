package main

import (
	"github.com/kormiltsev/filereaders/internal/app"
)

func main() {
	// // settings json
	// s := app.Settings
	// //err := app.UploadSettingsJson(&s)
	// err := app.CreateSettingsJson(&s)
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println("ok")
	// PG

	// files

	a := app.MiWatchSleep
	a.Do()

}
