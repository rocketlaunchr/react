// Copyright 2018 PJ Engineering and Business Solutions Pty. Ltd. All rights reserved.

package react

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

var (
	// React points to the React library. Change it
	// if it is not in your global namespace.
	//
	// See: https://www.npmjs.com/package/react

	React = js.Global.Get("React")

	// ReactDOM points to the ReactDOM library. Change it
	// if it is not in your global namespace.
	//
	// See: https://www.npmjs.com/package/react-dom
	ReactDOM = js.Global.Get("ReactDOM")

	// CreateReactClass points to create-react-class module.
	//
	// See: https://www.npmjs.com/package/create-react-class
	CreateReactClass = js.Global
)

// Render will render component to the specified target dom element.
func Render(element *js.Object, domTarget dom.Element, callback ...func()) *js.Object {
	if len(callback) > 0 && callback[0] != nil {
		return ReactDOM.Call("render", element, domTarget, callback[0])
	}
	return ReactDOM.Call("render", element, domTarget)
}
