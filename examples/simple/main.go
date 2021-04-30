package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

func main() {
	vecty.SetTitle("State example - bd")
	vecty.RenderBody(&PageView{})
}

type PageView struct {
	vecty.Core
}

func (p PageView) Render() vecty.ComponentOrHTML {

	counterone := NewCounter(myGlobalInt)
	countertwo := NewCounter(myGlobalInt)
	countertre := NewCounter(myOtherGlobalInt)
	counterfou := NewCounter(myOtherGlobalInt)

	return elem.Body(
		elem.Div(
			vecty.Text("The following two buttons share a counter."),
		),
		counterone,
		countertwo,
		elem.Div(
			vecty.Text("The following two buttons share a different counter."),
		),
		countertre,
		counterfou,
	)
}

