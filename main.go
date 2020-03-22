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

	// go goCount()

	vecty.RenderBody(page)
}

func goCount() {
	for {
		time.Sleep(1 * time.Second)
		page.Count++
	}
}

// PageView is our main page component.
type PageView struct {
	vecty.Core

	Count int
	Title string

	areaText string
	Area     *vecty.HTML
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
	p.areaText = p.Area.Node().Get("value").String()
	vecty.Rerender(p)
}

func (p *PageView) OnAreaChanged(event *vecty.Event) {
	p.areaText = event.Target.Get("value").String()
	vecty.Rerender(p)
}

// Render implements the vecty.Component interface.
func (p *PageView) Render() vecty.ComponentOrHTML {

	p.Count++

	go func() {
		for {
			time.Sleep(1 * time.Second)
			p.areaText = time.Now().String()
			vecty.Rerender(p)
		}
	}()

	p.Area = elem.TextArea(
		vecty.Markup(
			prop.Value(p.areaText),
			// prop.For(p.areaText),
			event.Change(p.OnAreaChanged),
		),
		vecty.Text("Text area"),
	)

	// p.Title = "This is Tittle"

	return elem.Body(

		elem.Label(
			vecty.Markup(
				prop.Value(string(p.Count)),
			),
			vecty.Text("Mark all as complete"),
		),

		elem.Break(),

		vecty.Text(p.Title),
		vecty.Text(strconv.Itoa(p.Count)),
		elem.Break(),

		p.Area,

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
				// prop.Value(p.areaText),
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
