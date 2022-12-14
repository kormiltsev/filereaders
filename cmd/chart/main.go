package main

import (
	"github.com/kormiltsev/filereaders/internal/chart"
	//"github.com/kormiltsev/filereaders/internal/storage"
)

func main() {
	chart := chart.NewSettings()
	chart.MakeChart()
}
