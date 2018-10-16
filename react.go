package main

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

var (
	// React points to the React library. Change it
	// if it is not in your global namespace.
	React = js.Global.Get("React")
	// ReactDOM points to the ReactDOM library. Change it
	// if it is not in your global namespace.
	ReactDOM = js.Global.Get("ReactDOM")
	// CreateReactClass points to create-react-class module.
	CreateReactClass = js.Global
)

// Render will render component to the specified target dom element.
func Render(element *js.Object, domTarget dom.Element, callback ...func()) *js.Object {
	if len(callback) > 0 && callback[0] != nil {
		return ReactDOM.Call("render", element, domTarget, callback[0])
	}
	return ReactDOM.Call("render", element, domTarget)
}
