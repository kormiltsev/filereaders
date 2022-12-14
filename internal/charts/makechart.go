package charts

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
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

var B = GetChart{
	FileName: "sample",
	Title:    "Testet",
	Subtitle: "noname",
	X:        []string{"1", "2", "3"},
	CatA:     []int{62, 67, 64},
	CatB:     []int{100, 104, 93},
	CatC:     []int{65, 65, 65},
}

func NewSettings() *GetChart {
	return &B
}
func generateBarItems(a []int) []opts.BarData {
	items := make([]opts.BarData, 0)
	for _, v := range a {
		items = append(items, opts.BarData{Value: v})
	}
	return items
}
func (cha *GetChart) MakeChart() {
	// create a new bar instance
	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    cha.Title,
		Subtitle: cha.Subtitle,
	}))

	// Put data into instance
	bar.SetXAxis(cha.X).
		//bar.SetXAxis(generDate()).
		AddSeries("Category A", generateBarItems(cha.CatA)).
		AddSeries("Category B", generateBarItems(cha.CatB))
	// Where the magic happens
	f, _ := os.Create(cha.FileName + ".html")
	bar.Render(f)
}
