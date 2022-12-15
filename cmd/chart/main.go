package main

import (
	"github.com/kormiltsev/filereaders/internal/chartoes"
	//"github.com/kormiltsev/filereaders/internal/storage"
)

func main() {
	chart := chartoes.NewSettings()
	chart.Export()
}
