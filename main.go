package main

import (
	"strconv"
	"time"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

var page = &PageView{}

func main() {
	vecty.SetTitle("Hello World!")
	// vecty.AddStylesheet("https://rawgit.com/tastejs/todomvc-common/master/base.css")
	// vecty.AddStylesheet("https://rawgit.com/tastejs/todomvc-app-css/master/index.css")

	// go goCount()

	page.Title = "Original"

	vecty.RenderBody(page)
}

func goCount() {
	for {
		time.Sleep(1 * time.Millisecond)
		page.Count++
	}
}

// PageView is our main page component.
type PageView struct {
	vecty.Core

	Count int
	Title string

	areaText string
	Area     vecty.HTML
}

func (p *PageView) OnRef(event *vecty.Event) {
	// .p.Title = "NEW"
	p.Title = "NEWWWW"
	vecty.Rerender(p)
	// p.Render()
}

func (p *PageView) OnToast(event *vecty.Event) {
	// .p.Title = "NEW"
	vecty.Rerender(p)
}

func (p *PageView) SetText(event *vecty.Event) {
	p.areaText = "sas"
	vecty.Rerender(p)
}

func (p *PageView) OnAreaChanged(event *vecty.Event) {
	p.areaText = event.Target.Get("value").String()
	vecty.Rerender(p)
}

// Render implements the vecty.Component interface.
func (p *PageView) Render() vecty.ComponentOrHTML {

	p.Count++

	// p.Title = "This is Tittle"

	return elem.Body(
		vecty.Text(p.Title),
		vecty.Text(strconv.Itoa(p.Count)),
		elem.Break(),

		elem.TextArea(
			vecty.Markup(
				// prop.Value(p.areaText),
				event.Change(p.OnAreaChanged),
			),
			vecty.Text("Text area"),
		),


		elem.Break(),

		elem.Input(
			vecty.Markup(
				vecty.Class("toggle-all"),
				prop.ID("toggle-all"),
				prop.Type(prop.TypeCheckbox),
			),
		),

		elem.Label(
			vecty.Markup(
				prop.For("toggle-all"),
			),
			vecty.Text("Mark all as complete"),
		),

		elem.Label(
			vecty.Markup(
				event.DoubleClick(p.OnToast),
			),
			vecty.Text("my label"),
		),

		elem.Break(),
		elem.Button(
			vecty.Markup(
				event.Click(p.OnRef),
			),
			vecty.Text("REF"),
		),

		elem.Break(),
		elem.Button(
			vecty.Markup(
				event.Click(p.OnToast),
			),
			vecty.Text("Toast"),
		),

		elem.Break(),
		elem.Label(
			vecty.Text(p.areaText),
		),
		elem.Break(),
		elem.Button(
			vecty.Markup(
				event.Click(p.SetText),
			),
			vecty.Text("SET TEXT"),
		),
	)
}
