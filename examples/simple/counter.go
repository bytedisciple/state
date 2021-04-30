//+build js

package main

import (
	"fmt"
	"github.com/bytedisciple/state"
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
	"github.com/hexops/vecty/style"
)

type Counter struct {
	vecty.Core
	val state.Int
}

func NewCounter(val state.Int) *Counter{
	c := Counter{
		val: val,
	}
	f1 := c.onMyGlobalIntChange
	val.Sub(&f1)
	return &c
}


func (c *Counter) onClick(event *vecty.Event) {
	c.val.Set(c.val.Get()+1)
}

// Render implements the vecty.Component interface.
func (c *Counter) Render() vecty.ComponentOrHTML {
	return elem.Button(
		vecty.Markup(
			prop.Href("#"),
			style.Height(style.Size("100")),
			style.Width(style.Size("100")),
			event.Click(c.onClick).PreventDefault(),
		),
		vecty.Text(fmt.Sprintf("Pressed %v times", c.val.Get())),
	)
}

func (c *Counter) onMyGlobalIntChange(oldValue, newValue int){
	vecty.Rerender(c)
}