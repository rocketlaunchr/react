// Copyright 2018-19 PJ Engineering and Business Solutions Pty. Ltd. All rights reserved.

package react

import (
	"github.com/gopherjs/gopherjs/js"
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
	CreateReactClass = js.Global.Get("createReactClass")

	// PureComponentMixin points to react-addons-pure-render-mixin module.
	// It is optional, but required if you want to create a PureComponent.
	//
	// Example:
	//
	//  pureDef := react.NewClassDef("Pure", react.PureComponentMixin)
	//
	// See: https://www.npmjs.com/package/react-addons-pure-render-mixin
	PureComponentMixin = js.Global.Get("PureRenderMixin")
)

// GetElementByID will return the first element with the specified id in the dom object.
// If no dom is provided, window.document will be used.
func GetElementByID(id string, dom ...*js.Object) *js.Object {
	if len(dom) > 0 {
		return dom[0].Call("getElementById", id)
	}
	return js.Global.Get("document").Call("getElementById", id)
}

// Render will render component to the specified target dom element.
func Render(element *js.Object, domTarget *js.Object, callback ...func()) *js.Object {
	if len(callback) > 0 && callback[0] != nil {
		return ReactDOM.Call("render", element, domTarget, callback[0])
	}
	return ReactDOM.Call("render", element, domTarget)
}
