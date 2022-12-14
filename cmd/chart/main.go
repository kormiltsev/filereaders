package main

import (
	"log"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/kormiltsev/filereaders/internal/storage"
)

// generate random data for bar chart
func generHR() []opts.BarData {
	items := make([]opts.BarData, 0)
	for _, v := range cat.MiwatchHrRow {
		log.Println(v)
		items = append(items, opts.BarData{Value: v.Heartrate})
	}
	return items
}

// func generDate() []string {
// 	dates := make([]string, 0)
// 	for _, v := range cat.MiwatchHrRow {
// 		dates = append(dates, strconv.Itoa(v.EventTimeInt))
// 	}
// 	return dates
// }

func main() {
	Update()
	// create a new bar instance
	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "My first bar chart generated by go-echarts",
		Subtitle: "It's extremely easy to use, right?",
	}))

	// Put data into instance
	bar.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
		//bar.SetXAxis(generDate()).
		AddSeries("Category A", generHR())
		//AddSeries("Category B", generateBarItems())
	// Where the magic happens
	f, _ := os.Create("bar.html")
	bar.Render(f)
}

func UploadCatalog() error {
	return storage.ConnectDB()
}

var cat = storage.PGbase{}

func Update() {
	storage.ConnectDB()
	cat = *storage.SendCatalog(10)
	//log.Println(cat)
}