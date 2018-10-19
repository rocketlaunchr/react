package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/rocketlaunchr/react"
	"github.com/rocketlaunchr/react/elements"
)

var EventsComponent *js.Object

type eventsState struct {
	// Since the counter can naturally have a value of zero, to make use of omitempty
	// (when setting state), we make the field a pointer.
	// That way, when we don't set the "Counter" field, it's 'zero-value' will be nil
	// instead of 0.
	Counter *int `react:"counter,omitempty"`
	// The sliderValue can't be 0 (the 'zero-value') due to the "min" attribute, so we
	// can keep it as a plain int.
	SliderValue int `react:"sliderValue,omitempty"`
}

func init() {

	eventsDef := react.NewClassDef("events")

	eventsDef["slider"] = react.CreateRef()

	eventsDef.GetInitialState(func(this *js.Object, props react.Map) interface{} {
		return eventsState{Counter: &[]int{0}[0], SliderValue: 0}
	})

	eventsDef.ComponentDidMount(func(this *js.Object, props, state react.Map, setState react.SetState) {

		slider := this.Get("slider").Get("current")
		val := slider.Get("value").Int()

		// Set state to initial slider value
		setState(eventsState{SliderValue: val})
	})

	// We can add to counter based on the state's "sliderValue" value directly, but for demonstration
	// purposes we'll pass it as an extra argument to the clickhandler event handler.
	eventsDef.SetMultiArgEventHandler("clickhandler", func(this *js.Object, arguments []*js.Object) func(this *js.Object, e *react.SyntheticEvent, props, state react.Map, setState react.SetState) {

		sliderValue := arguments[0].Int()

		return func(this *js.Object, e *react.SyntheticEvent, props, state react.Map, setState react.SetState) {

			var eState eventsState
			react.HydrateState(this, &eState)

			// Update counter
			setState(func(props, state react.Map) interface{} {
				*eState.Counter = *eState.Counter + sliderValue
				return eState
			})
		}

	})

	eventsDef.SetEventHandler("change", func(this *js.Object, e *react.SyntheticEvent, props, state react.Map, setState react.SetState) {

		slider := this.Get("slider").Get("current")
		val := slider.Get("value").Int()

		setState(eventsState{SliderValue: val})
	})

	eventsDef.Render(func(this *js.Object, props, state react.Map) interface{} {

		inputProps := &elements.InputProps{
			Type:     "range",
			Min:      "1",
			Max:      "5",
			OnChange: this.Get("change"),
			Ref:      this.Get("slider"),
		}

		sliderValue := state("sliderValue")

		buttonProps := &elements.ButtonProps{
			OnClick: this.Get("clickhandler").Invoke(sliderValue.Int()),
		}

		buttonText := "Add + " + state("sliderValue").String()

		return elements.Div(nil,
			elements.H3(&elements.H3Props{Style: &elements.Styles{TextAlign: "center"}}, "Counter: "+state("counter").String()),
			elements.Input(inputProps),
			elements.Button(buttonProps, buttonText),
		)

	})

	EventsComponent = react.CreateClass(eventsDef)
}
