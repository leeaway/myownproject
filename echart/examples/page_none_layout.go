package examples

import (
	"github.com/go-echarts/go-echarts/v2/components"
	"net/http"
)

type PageNoneLayoutExamples struct{}

func (PageNoneLayoutExamples) Examples(w http.ResponseWriter, _ *http.Request) {
	page := genPages()
	page.SetLayout(components.PageNoneLayout)
	page.Render(w)
}
