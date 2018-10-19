package main

import (
	"strconv"
	"time"

	"github.com/gopherjs/gopherjs/js"
	"github.com/rocketlaunchr/react"
	"github.com/rocketlaunchr/react/elements"
)

var TimerComponent *js.Object

type TimerProps struct {
	// We will pass in the time when the component was first instantiated.
	StartTime int64 `react:"start"`
}

type TimerState struct {
	// Elapsed will record how much time has passed since instantiation.
	Elapsed int64 `react:"elapsed"`
}

func init() {

	timerDef := react.NewClassDef("timer")

	timerDef.GetInitialState(func(this *js.Object, props react.Map) interface{} {

		// Using props in here is considered bad practice. It is only for example purposes.
		var tProps TimerProps
		react.HydrateProps(this, &tProps)

		return TimerState{
			Elapsed: time.Now().Unix() - tProps.StartTime,
		}
	})

	timerDef.ComponentDidMount(func(this *js.Object, props, state react.Map, setState react.SetState) {
		// Create a js timer that continually calls this.tick()
		this.Set("timer", react.JSFn("setInterval", this.Get("tick"), 1000))
	})

	timerDef.SetMethod("tick", func(this *js.Object, props, state react.Map, setState react.SetState, arguments []*js.Object) interface{} {

		var tProps TimerProps
		react.HydrateProps(this, &tProps)

		elapsed := time.Now().Unix() - tProps.StartTime

		// Update state
		setState(func(props, state react.Map) interface{} {
			tState := TimerState{
				Elapsed: elapsed,
			}
			return tState
		}, func() {
			println("state.elapsed updated to:", state("elapsed"))
		})

		return nil
	})

	timerDef.Render(func(this *js.Object, props, state react.Map) interface{} {

		var tState TimerState
		react.HydrateState(this, &tState)

		text := "Uptime counter:" + strconv.Itoa(int(tState.Elapsed))

		return elements.H3(&elements.H3Props{Style: &elements.Styles{TextAlign: "center"}}, text)
	})

	TimerComponent = react.CreateClass(timerDef)
}
