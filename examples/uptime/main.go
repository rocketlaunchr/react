package main

import (
	"github.com/rocketlaunchr/react"
	"honnef.co/go/js/dom"
)

var document = dom.GetWindow().Document()

func main() {
	domTarget := document.GetElementByID("app")

	title := "UPTIME TIMER"

	react.Render(react.JSX(ContainerComponent, &ContainerProps{Title: title}), domTarget)

}
