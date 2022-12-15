package main

import (
	"net/http"

	"github.com/kormiltsev/filereaders/internal/chartoes"
	//"github.com/kormiltsev/filereaders/internal/storage"
)

func main() {
	chart := chartoes.NewSettings()
	chart.Export()

	http.HandleFunc("/", chartoes.httpCharts(chart))
	http.ListenAndServe(":8081", nil)
}
