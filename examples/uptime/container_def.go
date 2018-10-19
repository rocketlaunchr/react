package main

import (
	"time"

	"github.com/gopherjs/gopherjs/js"
	"github.com/rocketlaunchr/react"
)

var ContainerComponent *js.Object

type ContainerProps struct {
	Title string `react:"title"`
}

func init() {

	containerDef := react.NewClassDef("container")

	containerDef.Render(func(this *js.Object, props, state react.Map) interface{} {
		title := props("title").String()

		return react.Fragment(nil,
			react.JSX(TitleComponent, &TitleProps{Title: title}),
			react.JSX(TimerComponent, &TimerProps{StartTime: time.Now().Unix()}),
		)

	})

	ContainerComponent = react.CreateClass(containerDef)
}
