package examples

import (
	"github.com/go-echarts/go-echarts/v2/components"
	"net/http"
)

func genPages() *components.Page {
	page := components.NewPage()
	page.AddCharts(
		bar3DAutoRotate(),
		gaugeBase(),
		mapShowLabel(),
		graphBase(),
		graphCircle(),
		geoBase(),
		geoGuangdong(),
		geoJiangxi(),
	)
	return page
}

type PageCenterLayoutExamples struct{}

func (PageCenterLayoutExamples) Examples(w http.ResponseWriter, _ *http.Request) {
	page := genPages()
	page.Render(w)
}
