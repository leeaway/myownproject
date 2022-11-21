package echart

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"math/rand"
	"net/http"
)

// generate random data for line chart
func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}

func Httpserver(w http.ResponseWriter, _ *http.Request) {
	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Line example in Westeros theme",
			Subtitle: "Line chart rendered by the http server this time",
		}))

	// Put data into instance
	line.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
		AddSeries("Category A", generateLineItems()).
		AddSeries("Category B", generateLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	line.Render(w)
}

func HttpserverPie(w http.ResponseWriter, _ *http.Request) {
	// create a new line instance
	pie := charts.NewPie()
	// set some global options like Title/Legend/ToolTip or anything else
	pie.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeShine}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Pie example",
			Subtitle: "Pie chart rendered by the http server this time",
		}))

	// Put data into instance
	var params []opts.PieData

	params = append(params, opts.PieData{
		Name:     fmt.Sprintf("Location CC"),
		Value:    16436,
		Selected: true,
		Label: &opts.Label{
			Show:     true,
			Color:    "auto",
			Position: "top",
		},
		ItemStyle: &opts.ItemStyle{
			Color:       "orange",
			BorderColor: "auto",
		},
		Tooltip: &opts.Tooltip{
			Show:      true,
			Trigger:   "item",
			TriggerOn: "mousemove|click",
			Formatter: "{c}",
			AxisPointer: &opts.AxisPointer{
				Snap: true,
			},
		},
	})

	params = append(params, opts.PieData{
		Name:     fmt.Sprintf("SKU CC"),
		Value:    12932,
		Selected: false,
		Label: &opts.Label{
			Show:     true,
			Color:    "auto",
			Position: "top",
		},
		ItemStyle: &opts.ItemStyle{
			Color:       "grey",
			BorderColor: "auto",
		},
		Tooltip: &opts.Tooltip{
			Show:      true,
			Trigger:   "item",
			TriggerOn: "mousemove|click",
			Formatter: "{c}",
			AxisPointer: &opts.AxisPointer{
				Snap: true,
			},
		},
	})

	params = append(params, opts.PieData{
		Name:     fmt.Sprintf("Movement CC"),
		Value:    1674,
		Selected: false,
		Label: &opts.Label{
			Show:     true,
			Color:    "auto",
			Position: "top",
		},
		ItemStyle: &opts.ItemStyle{
			Color:       "green",
			BorderColor: "auto",
		},
		Tooltip: &opts.Tooltip{
			Show:      true,
			Trigger:   "item",
			TriggerOn: "mousemove|click",
			Formatter: "{c}",
			AxisPointer: &opts.AxisPointer{
				Snap: true,
			},
		},
	})

	pie.AddSeries("test pie", params)
	pie.Render(w)
}
