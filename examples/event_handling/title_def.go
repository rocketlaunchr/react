package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/rocketlaunchr/react"
	"github.com/rocketlaunchr/react/elements"
)

// TitleComponent is a react component.
var TitleComponent *js.Object

// TitleProps are the props for TitleComponent.
type TitleProps struct {
	Title string `react:"title"`
}

func init() {

	titleDef := react.NewClassDef("Title")

	titleDef.Render(func(this *js.Object, props, state react.Map) interface{} {

		var tProps TitleProps
		react.HydrateProps(this, &tProps)

		title := tProps.Title

		return elements.H1(&elements.H1Props{Style: &elements.Styles{TextAlign: "center"}}, title)
	})

	TitleComponent = react.CreateClass(titleDef)
}
