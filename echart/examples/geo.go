package examples

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"math/rand"
	"net/http"
)

var geoData = []opts.GeoData{
	{Name: "SG", Value: []float64{103.45, 1.22, float64(rand.Intn(100))}},
	{Name: "PH", Value: []float64{121, 14.37, float64(rand.Intn(100))}},
	{Name: "MY", Value: []float64{101.42, 3.08, float64(rand.Intn(100))}},
	{Name: "ID", Value: []float64{106.45, 6.08, float64(rand.Intn(100))}},
	{Name: "TW", Value: []float64{121.38, 25.03, float64(rand.Intn(100))}},
	{Name: "VN", Value: []float64{105.53, 21.01, float64(rand.Intn(100))}},
	{Name: "CN", Value: []float64{116.23, 39.55, float64(rand.Intn(100))}},
	{Name: "TH", Value: []float64{100.29, 13.5, float64(rand.Intn(100))}},
	{Name: "MX", Value: []float64{-99.09, 19.28, float64(rand.Intn(100))}},
}

var guangdongData = []opts.GeoData{
	{Name: "汕头", Value: []float64{116.69, 23.39, float64(rand.Intn(100))}},
	{Name: "深圳", Value: []float64{114.07, 22.62, float64(rand.Intn(100))}},
	{Name: "广州", Value: []float64{113.23, 23.16, float64(rand.Intn(100))}},
}

func geoBase() *charts.Geo {
	geo := charts.NewGeo()
	geo.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic geo example"}),
		charts.WithGeoComponentOpts(opts.GeoComponent{
			Map:       "world",
			ItemStyle: &opts.ItemStyle{Color: "#006666"},
		}),
	)

	geo.AddSeries("geo", types.ChartEffectScatter, geoData,
		charts.WithRippleEffectOpts(opts.RippleEffect{
			Period:    4,
			Scale:     6,
			BrushType: "stroke",
		}),
	)
	return geo
}

func geoGuangdong() *charts.Geo {
	geo := charts.NewGeo()
	geo.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Guangdong province"}),
		charts.WithGeoComponentOpts(opts.GeoComponent{
			Map: "广东",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			InRange: &opts.VisualMapInRange{
				Color: []string{"#50a3ba", "#eac736", "#d94e5d"},
			},
		}),
	)

	geo.AddSeries("geo", types.ChartScatter, guangdongData)
	return geo
}

func geoJiangxi() *charts.Geo {
	geo := charts.NewGeo()
	geo.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Jiangxi province"}),
		charts.WithGeoComponentOpts(opts.GeoComponent{
			Map: "江西",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			InRange: &opts.VisualMapInRange{
				Color: []string{"#50a3ba", "#eac736", "#d94e5d"},
			},
		}),
	)

	geo.AddSeries("geo", types.ChartScatter, guangdongData)
	return geo
}

type GeoExamples struct{}

func (GeoExamples) Examples(w http.ResponseWriter, _ *http.Request) {
	page := components.NewPage()
	page.AddCharts(
		geoBase(),
		geoGuangdong(),
		geoJiangxi(),
	)
	page.Render(w)
}
