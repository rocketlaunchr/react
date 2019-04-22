// Copyright 2018 PJ Engineering and Business Solutions Pty. Ltd. All rights reserved.

package react

import (
	"github.com/gopherjs/gopherjs/js"
)

// SyntheticEvent represents a SyntheticEvent.
//
// See: https://reactjs.org/docs/events.html#overview
type SyntheticEvent struct {
	// O represents the original React SyntheticEvent.
	O *js.Object
}

// Bubbles
//
// See: https://reactjs.org/docs/events.html#overview
func (s *SyntheticEvent) Bubbles() bool {
	return s.O.Get("bubbles").Bool()
}

// Cancelable
//
// See: https://reactjs.org/docs/events.html#overview
func (s *SyntheticEvent) Cancelable() bool {
	return s.O.Get("cancelable").Bool()
}

// CurrentTarget
//
// See: https://reactjs.org/docs/events.html#overview
//
// Example:
//
//  import "honnef.co/go/js/dom"
//
//  dom.WrapHTMLElement(e.CurrentTarget())
func (s *SyntheticEvent) CurrentTarget() *js.Object {
	return s.O.Get("currentTarget")
}

// DefaultPrevented
//
// See: https://reactjs.org/docs/events.html#overview
func (s *SyntheticEvent) DefaultPrevented() bool {
	return s.O.Get("defaultPrevented").Bool()
}

// EventPhase
//
// See: https://reactjs.org/docs/events.html#overview
func (s *SyntheticEvent) EventPhase() int {
	return s.O.Get("eventPhase").Int()
}

// IsTrusted
//
// See: https://reactjs.org/docs/events.html#overview
func (s *SyntheticEvent) IsTrusted() bool {
	return s.O.Get("isTrusted").Bool()
}

// NativeEvents
//
// See: https://reactjs.org/docs/events.html#overview
//
// Example:
//
//  import "honnef.co/go/js/dom"
//
//  dom.WrapEvent(e.NativeEvent())
func (s *SyntheticEvent) NativeEvent() *js.Object {
	return s.O.Get("nativeEvent")
}

// PreventDefault
//
// See: https://reactjs.org/docs/events.html#overview
func (s *SyntheticEvent) PreventDefault() {
	s.O.Call("preventDefault")
}

// IsDefaultPrevented
// See: https://reactjs.org/docs/events.html#overview
func (s *SyntheticEvent) IsDefaultPrevented() bool {
	return s.O.Call("isDefaultPrevented").Bool()
}

// StopPropagation
//
// See: https://reactjs.org/docs/events.html#overview
func (s *SyntheticEvent) StopPropagation() {
	s.O.Call("stopPropagation")
}

// IsPropagationStopped
//
// See: https://reactjs.org/docs/events.html#overview
func (s *SyntheticEvent) IsPropagationStopped() bool {
	return s.O.Call("isPropagationStopped").Bool()
}

// Target
//
// See: https://reactjs.org/docs/events.html#overview
//
// Example:
//
//  import "honnef.co/go/js/dom"
//
//  dom.WrapHTMLElement(e.Target())
func (s *SyntheticEvent) Target() *js.Object {
	return s.O.Get("target")
}

// TimeStamp
//
// See: https://reactjs.org/docs/events.html#overview
func (s *SyntheticEvent) TimeStamp() float64 {
	return s.O.Get("timeStamp").Float()
}

// Type
//
// See: https://reactjs.org/docs/events.html#overview
func (s *SyntheticEvent) Type() string {
	return s.O.Get("type").String()
}

// Persist is used if you want to access properties in an asynchronous way.
//
// See: https://reactjs.org/docs/events.html#event-pooling
func (s *SyntheticEvent) Persist() *SyntheticEvent {
	p := s.O.Call("persist")
	return &SyntheticEvent{p}
}

// SetEventHandler allows a custom event handler to be attached.
// By passing nil for f, the handler can also be detached (cleared).
//
// It can be used like this: "onClick": this.Get("clickhandler")
//
func (def ClassDef) SetEventHandler(name string, f func(this *js.Object, e *SyntheticEvent, props, state Map, setState SetState)) {

	h := func(this *js.Object, props, state Map, setState SetState, arguments []*js.Object) interface{} {
		syntheticEvent := &SyntheticEvent{arguments[0]}
		f(this, syntheticEvent, props, state, setState)
		return nil
	}

	def.setMethod(false, name, h)
}

// SetMultiArgEventHandler allows for you to pass custom arguments to a custom
// event handler. By passing nil for f, the handler can also be detached (cleared).
//
// It can be used like this: "onClick": this.Get("clickhandler").Invoke(5)
//
// See: https://reactjs.org/docs/handling-events.html#passing-arguments-to-event-handlers
func (def ClassDef) SetMultiArgEventHandler(name string, f func(this *js.Object, arguments []*js.Object) func(this *js.Object, e *SyntheticEvent, props, state Map, setState SetState)) {

	if f == nil {
		// Clear handler
		delete(def, name)
		return
	}

	if name == "statics" {
		panic("can't have function name called 'statics'")
	}

	x := func(this *js.Object, arguments []*js.Object) interface{} {

		props := func(key string) *js.Object {
			return this.Get("props").Get(key)
		}

		state := func(key string) *js.Object {
			return this.Get("state").Get(key)
		}

		setState := func(updater interface{}, callback ...func()) {

			if updater == nil {
				return
			}

			if len(callback) > 0 && callback[0] != nil {
				switch updater := updater.(type) {
				case func(props, state Map) interface{}:
					this.Call("setState", func(state *js.Object, props *js.Object) interface{} {
						return SToMap(updater(func(key string) *js.Object {
							return props.Get(key)
						}, func(key string) *js.Object {
							return state.Get(key)
						}))
					}, callback[0])
				case UpdaterFunc:
					this.Call("setState", func(state *js.Object, props *js.Object) interface{} {
						return SToMap(updater(func(key string) *js.Object {
							return props.Get(key)
						}, func(key string) *js.Object {
							return state.Get(key)
						}))
					}, callback[0])
				default:
					this.Call("setState", SToMap(updater), callback[0])
				}
			} else {
				switch updater := updater.(type) {
				case func(props, state Map) interface{}:
					this.Call("setState", func(state *js.Object, props *js.Object) interface{} {
						return SToMap(updater(func(key string) *js.Object {
							return props.Get(key)
						}, func(key string) *js.Object {
							return state.Get(key)
						}))
					})
				case UpdaterFunc:
					this.Call("setState", func(state *js.Object, props *js.Object) interface{} {
						return SToMap(updater(func(key string) *js.Object {
							return props.Get(key)
						}, func(key string) *js.Object {
							return state.Get(key)
						}))
					})
				default:
					this.Call("setState", SToMap(updater))
				}
			}
		}

		z := f(this, arguments)

		return func(e *js.Object) {
			syntheticEvent := &SyntheticEvent{e}
			z(this, syntheticEvent, props, state, setState)
		}
	}

	def[name] = js.MakeFunc(x)
}
