package chartoes

import (
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	//"github.com/kormiltsev/filereaders/internal/storage"
)

type GetChart struct {
	FileName string
	Title    string
	Subtitle string
	X        []string
	CatA     []int
	CatB     []int
	CatC     []int
}
type MapOfCharts struct {
	FileName string
	HR       *GetChart
	Sleep    *GetChart
}

var B = MapOfCharts{
	"sample.html",
	&heartrateChart,
	&heartrateChart,
}
var heartrateChart = GetChart{
	FileName: "sample",
	Title:    "Testet",
	Subtitle: "noname",
	X:        []string{"1", "2", "3", "4"},
	CatA:     []int{62, 67, 64, 62},
	CatB:     []int{100, 104, 0, 93},
	CatC:     []int{65, 65, 65, 65},
}

func NewSettings() *MapOfCharts {
	return &B
}
func generateBarItems(a []int) []opts.BarData {
	items := make([]opts.BarData, 0)
	for _, v := range a {
		items = append(items, opts.BarData{Value: v})
	}
	return items
}
func (cha *GetChart) MakeChartBar() *charts.Bar {
	// create a new bar instance
	chbar := charts.NewBar()

	// set some global options like Title/Legend/ToolTip or anything else
	chbar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    cha.Title,
		Subtitle: cha.Subtitle,
	}))

	// Put data into instance
	chbar.SetXAxis(cha.X).
		//bar.SetXAxis(generDate()).
		AddSeries("Category A", generateBarItems(cha.CatA)).
		AddSeries("Category B", generateBarItems(cha.CatB))
	// Where the magic happens
	// f, _ := os.Create(cha.FileName + ".html")
	// chbar.Render(f)
	return chbar
}

func generateLineItems(a []int) []opts.LineData {
	items := make([]opts.LineData, 0)
	for _, v := range a {
		items = append(items, opts.LineData{Value: v})
	}
	return items
}

func (cha *GetChart) MakeChartLine() *charts.Line {
	// create a new bar instance
	chline := charts.NewLine()
	chline.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		// set some global options like Title/Legend/ToolTip or anything else
		charts.WithTitleOpts(opts.Title{
			Title:    cha.Title,
			Subtitle: cha.Subtitle,
			Link:     "https://ya.ru",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "Bit per min",
			Min:  60,
			SplitLine: &opts.SplitLine{
				Show: false,
			},
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "Elements",
		}),
	)

	// Put data into instance
	chline.SetXAxis(cha.X).
		//bar.SetXAxis(generDate()).
		AddSeries("Category A", generateLineItems(cha.CatA)).
		AddSeries("Category B", generateLineItems(cha.CatB)).
		AddSeries("Category C", generateLineItems(cha.CatC)).
		SetSeriesOptions(
			charts.WithLineChartOpts(opts.LineChart{Smooth: true}), //, ConnectNulls: false}),
			charts.WithLabelOpts(opts.Label{Show: true}),
		)
	// Where the magic happens
	// f, _ := os.Create(cha.FileName + ".html")
	// chline.Render(f)
	return chline
}

func (c *MapOfCharts) Export() {
	page := components.NewPage()
	page.AddCharts(
		B.HR.MakeChartLine(),
		B.Sleep.MakeChartBar(),
		// lineShowLabel(),
		// lineMarkPoint(),
		// lineSplitLine(),
		// lineStep(),
		// lineSmooth(),
		// lineArea(),
		// lineSmoothArea(),
		// lineOverlap(),
		// lineMulti(),
		// lineDemo(),
	)
	f, err := os.Create(c.FileName)
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
