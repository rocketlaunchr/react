package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/rocketlaunchr/react"
	"honnef.co/go/js/dom"
)

var document = dom.GetWindow().Document()

func main() {
	domTarget := document.GetElementByID("app")

	// An example using Functional Components
	// See: https://reactjs.org/docs/components-and-props.html
	type ContainerProps struct {
		Title string `react:"title"`
	}

	title := "EVENTS HANDLING"

	containerComponent := func(props *js.Object) *js.Object {
		title := props.Get("title").String()

		return react.Fragment(nil,
			react.JSX(TitleComponent, &TitleProps{Title: title}),
			react.JSX(EventsComponent, nil),
		)
	}

	react.Render(react.JSX(containerComponent, &ContainerProps{Title: title}), domTarget)
}
