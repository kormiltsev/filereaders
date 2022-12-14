package main

import (
	"log"

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

	if app.Start() != nil {
		log.Println("STOPPED")
		return
	}
	a := app.MiWatchSleep
	a.Do()

	b := app.MiWatchHr
	b.Do()

	app.Finish()
}
