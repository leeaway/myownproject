package examples

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"math/rand"
	"net/http"
)

func gaugeBase() *charts.Gauge {
	gauge := charts.NewGauge()
	gauge.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic Gauge example"}),
	)

	gauge.AddSeries("ProjectA", []opts.GaugeData{{Name: "Work progress", Value: rand.Intn(50)}})
	return gauge
}

func gaugeTimer() *charts.Gauge {
	gauge := charts.NewGauge()
	gauge.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "javascript timer"}),
	)

	gauge.AddSeries("ProjectB", []opts.GaugeData{{Name: "Work progress", Value: rand.Intn(50)}})

	fn := fmt.Sprintf(`setInterval(function () {
			option_%s.series[0].data[0].value = (Math.random() * 100).toFixed(2) - 0;
			goecharts_%s.setOption(option_%s, true);
		}, 2000);`, gauge.ChartID, gauge.ChartID, gauge.ChartID)
	gauge.AddJSFuncs(fn)
	return gauge
}

type GaugeExamples struct{}

func (GaugeExamples) Examples(w http.ResponseWriter, _ *http.Request) {
	page := components.NewPage()
	page.AddCharts(
		gaugeBase(),
		gaugeTimer(),
	)
	page.Render(w)
}
