package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/rocketlaunchr/react"
	"github.com/rocketlaunchr/react/elements"
)

var TitleComponent *js.Object

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
