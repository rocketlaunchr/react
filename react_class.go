// Copyright 2018 PJ Engineering and Business Solutions Pty. Ltd. All rights reserved.

package react

import (
	"github.com/gopherjs/gopherjs/js"
)

// Map is a convenience method that can be used to access fields in a
// js object.
type Map func(key string) *js.Object

// UpdaterFunc is the first argument for SetState function.
//
// See: https://reactjs.org/docs/react-component.html#setstate
type UpdaterFunc func(props, state Map) interface{}

// SetState is used to asynchronously update the state.
// If the new state is dependent on the current props or state,
// updater must be of type UpdaterFunc.
//
// See: https://reactjs.org/docs/react-component.html#setstate
type SetState func(updater interface{}, callback ...func())

// ForceUpdate will force a rerender of the component.
//
// See: https://reactjs.org/docs/react-component.html#forceupdate
func ForceUpdate(this *js.Object, callback ...func()) {
	if len(callback) > 0 && callback[0] != nil {
		this.Call("forceUpdate", callback[0])
	} else {
		this.Call("forceUpdate")
	}
}

// ClassDef is used to create custom React components.
type ClassDef map[string]interface{}

// NewClassDef will create an empty class definition which can immediately be used
// to create a React component. displayName is the text that is shown in Chrome's
// React Developer Tools.
func NewClassDef(displayName string) ClassDef {
	def := ClassDef{
		render: js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
			return nil
		}),
	}
	def["displayName"] = displayName
	return def
}

func (def ClassDef) setMethod(static bool, name string, f func(this *js.Object, props, state Map, setState SetState, arguments []*js.Object) interface{}) {

	const statics = "statics"

	if f == nil {
		// Clear method
		if static {
			if _, exists := def[statics]; exists {
				switch s := def[statics].(type) {
				case (map[string]interface{}):
					delete(s, name)
				default:

				}
			}
		} else {
			delete(def, name)
		}
		return
	}

	if !static && name == statics {
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

		return f(this, props, state, setState, arguments)
	}

	if static {
		def[statics] = map[string]interface{}{
			name: js.MakeFunc(x),
		}
	} else {
		def[name] = js.MakeFunc(x)
	}
}

// SetMethod allows a custom method to be attached.
// By passing nil for f, the method can also be detached (cleared).
func (def ClassDef) SetMethod(name string, f func(this *js.Object, props, state Map, setState SetState, arguments []*js.Object) interface{}) {
	def.setMethod(false, name, f)
}

// CreateClass is used to create a react component.
//
// See: https://reactjs.org/docs/react-without-es6.html
func CreateClass(def ClassDef) *js.Object {
	return CreateReactClass.Call("createReactClass", def)
}
