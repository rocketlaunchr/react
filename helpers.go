// Copyright 2018 PJ Engineering and Business Solutions Pty. Ltd. All rights reserved.

package react

import (
	"github.com/gopherjs/gopherjs/js"
)

// Fragment is used to group a list of children
// without adding extra nodes to the DOM.
//
// See: https://reactjs.org/docs/fragments.html
func Fragment(key *string, children ...interface{}) *js.Object {
	props := map[string]interface{}{}
	if key != nil {
		props["key"] = *key
	}
	return JSX(React.Get("Fragment"), props, children...)
}

// JSX is used to create an Element.
func JSX(component interface{}, props interface{}, children ...interface{}) *js.Object {

	args := []interface{}{
		component,
		SToMap(props),
	}
	if len(children) > 0 {
		args = append(args, children...)
	}

	return React.Call("createElement", args...)
}

// JSFn is a convenience function used to call javascript functions that are
// part of the standard library.
func JSFn(name string, args ...interface{}) *js.Object {
	return js.Global.Call(name, args...)
}

// CreateRef will create a Ref.
//
// See: https://reactjs.org/docs/refs-and-the-dom.html
func CreateRef() *js.Object {
	return React.Call("createRef")
}

// ForwardRef will forward a Ref to child components.
//
// See: https://reactjs.org/docs/forwarding-refs.html
func ForwardRef(component interface{}) *js.Object {
	return React.Call("forwardRef", func(props *js.Object, ref *js.Object) *js.Object {
		props.Set("ref", ref)

		n := React.Get("Children").Call("count", props.Get("children")).Int()
		switch n {
		case 0:
			return JSX(component, props)
		case 1:
			return JSX(component, props, props.Get("children"))
		default:
			children := []interface{}{}
			for i := 0; i < n; i++ {
				children = append(children, props.Get("children").Index(i))
			}
			return JSX(component, props, children...)
		}
	})
}

// CreateContext is used when you want to pass data to a deeply
// embedded child component without using props.
//
// See: https://reactjs.org/docs/context.html#reactcreatecontext
func CreateContext(defaultValue ...interface{}) (Provider *js.Object, Consumer *js.Object) {

	var res *js.Object

	if len(defaultValue) > 0 {
		res = React.Call("createContext", defaultValue[0])
	} else {
		res = React.Call("createContext")
	}

	return res.Get("Provider"), res.Get("Consumer")
}

// CloneElement is used to clone and return a new React Element.
//
// See: https://reactjs.org/docs/react-api.html#cloneelement
func CloneElement(element interface{}, props interface{}, children ...interface{}) *js.Object {

	args := []interface{}{
		element,
		SToMap(props),
	}
	if len(children) > 0 {
		args = append(args, children...)
	}

	return React.Call("cloneElement", args...)
}
