package main

import (
	"github.com/rocketlaunchr/react"
)

func main() {
	domTarget := react.CreateRoot(react.GetElementByID("app"))

	title := "UPTIME TIMER"

	react.Render(react.JSX(ContainerComponent, &ContainerProps{Title: title}), domTarget)
}
