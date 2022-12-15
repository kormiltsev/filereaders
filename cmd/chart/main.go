package main

import (
	"log"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/kormiltsev/filereaders/internal/chartoes"
	//"github.com/kormiltsev/filereaders/internal/storage"
)

func httpCharts(w http.ResponseWriter, _ *http.Request) {
	B := chartoes.NewSettings()
	page := components.NewPage()
	page.AddCharts(
		B.HR.MakeChartLine(),
		B.Sleep.MakeChartBar(),
	)
	page.Render(w)
}

func main() {
	chart := chartoes.NewSettings()
	chart.Export()

	http.HandleFunc("/", httpCharts)
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
