package examples

import (
	"github.com/go-echarts/go-echarts/v2/components"
	"net/http"
)

type PageFlexLayoutExamples struct{}

func (PageFlexLayoutExamples) Examples(w http.ResponseWriter, _ *http.Request) {
	page := genPages()
	page.SetLayout(components.PageFlexLayout)
	page.Render(w)
}
