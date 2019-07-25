package main

import (
	"time"

	"github.com/gopherjs/gopherjs/js"
	"github.com/rocketlaunchr/react"
)

// ContainerComponent is a react component.
var ContainerComponent *js.Object

// ContainerProps are the props for ContainerComponent.
type ContainerProps struct {
	Title string `react:"title"`
}

func init() {

	containerDef := react.NewClassDef("Container")

	containerDef.Render(func(this *js.Object, props, state react.Map) interface{} {
		title := props("title").String()

		return react.Fragment(nil,
			react.JSX(TitleComponent, &TitleProps{Title: title}),
			react.JSX(TimerComponent, &TimerProps{StartTime: time.Now().Unix()}),
		)

	})

	ContainerComponent = react.CreateClass(containerDef)
}
