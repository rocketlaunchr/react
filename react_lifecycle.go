// Copyright 2018 PJ Engineering and Business Solutions Pty. Ltd. All rights reserved.

package react

import (
	"github.com/gopherjs/gopherjs/js"
)

const (
	// TODO: add type checking props
	getDefaultProps = "getDefaultProps"

	// Mounting
	getInitialState          = "getInitialState"
	getDerivedStateFromProps = "getDerivedStateFromProps"
	render                   = "render"
	componentDidMount        = "componentDidMount"

	// Updating
	shouldComponentUpdate   = "shouldComponentUpdate"
	getSnapshotBeforeUpdate = "getSnapshotBeforeUpdate"
	componentDidUpdate      = "componentDidUpdate"

	// Unmounting
	componentWillUnmount = "componentWillUnmount"

	// Error-handling
	componentDidCatch = "componentDidCatch"
)

// GetDefaultProps sets the getDefaultProps method.
func (def ClassDef) GetDefaultProps(f func(this *js.Object) interface{}) {
	def.SetMethod(getDefaultProps, func(this *js.Object, props, state Map, setState SetState, arguments []*js.Object) interface{} {
		return SToMap(f(this))
	})
}

// GetInitialState sets the getInitialState method.
// Note: It is usually not recommended to use the props when setting the initial state.
func (def ClassDef) GetInitialState(f func(this *js.Object, props Map) interface{}) {
	def.SetMethod(getInitialState, func(this *js.Object, props, state Map, setState SetState, arguments []*js.Object) interface{} {
		return SToMap(f(this, props))
	})
}

// GetDerivedStateFromProps sets the getDerivedStateFromProps class method.
//
// See: https://reactjs.org/blog/2018/06/07/you-probably-dont-need-derived-state.html
func (def ClassDef) GetDerivedStateFromProps(f func(nextProps, prevState Map) interface{}) {

	def.setMethod(true, getDerivedStateFromProps, func(this *js.Object, props, state Map, setState SetState, arguments []*js.Object) interface{} {

		nextProps := func(key string) *js.Object {
			return arguments[0].Get(key)
		}
		prevState := func(key string) *js.Object {
			return arguments[1].Get(key)
		}

		return SToMap(f(nextProps, prevState))
	})
}

// ComponentDidMount sets the componentDidMount method.
//
// See: https://reactjs.org/docs/react-component.html#componentdidmount
func (def ClassDef) ComponentDidMount(f func(this *js.Object, props, state Map, setState SetState)) {
	def.SetMethod(componentDidMount, func(this *js.Object, props, state Map, setState SetState, arguments []*js.Object) interface{} {
		f(this, props, state, setState)
		return nil
	})
}

// ComponentWillUnmount sets the componentWillUnmount method.
//
// See: https://reactjs.org/docs/react-component.html#componentwillunmount
func (def ClassDef) ComponentWillUnmount(f func(this *js.Object, props, state Map)) {
	def.SetMethod(componentWillUnmount, func(this *js.Object, props, state Map, setState SetState, arguments []*js.Object) interface{} {
		f(this, props, state)
		return nil
	})
}

// ShouldComponentUpdate sets the shouldComponentUpdate method.
//
// See: https://reactjs.org/docs/react-component.html#shouldcomponentupdate
func (def ClassDef) ShouldComponentUpdate(f func(this *js.Object, props, nextProps, state, nextState Map) bool) {
	def.SetMethod(shouldComponentUpdate, func(this *js.Object, props, state Map, setState SetState, arguments []*js.Object) interface{} {
		nextProps := func(key string) *js.Object {
			return arguments[0].Get(key)
		}
		nextState := func(key string) *js.Object {
			return arguments[1].Get(key)
		}
		return f(this, props, nextProps, state, nextState)
	})
}

// GetSnapshotBeforeUpdate sets the getSnapshotBeforeUpdate method.
//
// See: https://reactjs.org/docs/react-component.html#getsnapshotbeforeupdate
func (def ClassDef) GetSnapshotBeforeUpdate(f func(this *js.Object, prevProps, props, prevState, state Map) interface{}) {
	def.SetMethod(getSnapshotBeforeUpdate, func(this *js.Object, props, state Map, setState SetState, arguments []*js.Object) interface{} {
		prevProps := func(key string) *js.Object {
			return arguments[0].Get(key)
		}
		prevState := func(key string) *js.Object {
			return arguments[1].Get(key)
		}

		ret := f(this, prevProps, props, prevState, state)
		if ret == nil {
			return nil
		} else if isStruct(ret) {
			return convertStruct(ret)
		} else {
			return ret
		}
	})
}

// ComponentDidUpdate sets the componentDidUpdate method.
//
// See: https://reactjs.org/docs/react-component.html#componentdidupdate
func (def ClassDef) ComponentDidUpdate(f func(this *js.Object, prevProps, props, prevState, state Map, setState SetState, snapshot *js.Object)) {
	def.SetMethod(componentDidUpdate, func(this *js.Object, props, state Map, setState SetState, arguments []*js.Object) interface{} {
		snapshot := arguments[2]
		prevProps := func(key string) *js.Object {
			return arguments[0].Get(key)
		}
		prevState := func(key string) *js.Object {
			return arguments[1].Get(key)
		}
		f(this, prevProps, props, prevState, state, setState, snapshot)
		return nil
	})
}

// Render sets the render method.
//
// See: https://reactjs.org/docs/react-component.html#render
func (def ClassDef) Render(f func(this *js.Object, props, state Map) interface{}) {
	def.SetMethod(render, func(this *js.Object, props, state Map, setState SetState, arguments []*js.Object) interface{} {
		return f(this, props, state)
	})
}

// ComponentDidCatch sets the componentDidCatch method.
//
// See: https://reactjs.org/docs/react-component.html#componentdidcatch
func (def ClassDef) ComponentDidCatch(f func(this *js.Object, err, info *js.Object, props, state Map, setState SetState)) {
	def.SetMethod(componentDidCatch, func(this *js.Object, props, state Map, setState SetState, arguments []*js.Object) interface{} {
		err := arguments[0]
		info := arguments[1]
		f(this, err, info, props, state, setState)
		return nil
	})
}
