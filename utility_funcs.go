// Copyright 2018 PJ Engineering and Business Solutions Pty. Ltd. All rights reserved.

package react

import (
	"strings"
)

// Set is used for conveniently dealing with
// data-* and aria-* attributes.
//
// See: https://reactjs.org/docs/dom-elements.html
type Set map[string]string

// Convert is used to transform a set of suffix attributes
// to the actual attributes by prefixing them with a base.
func (s Set) Convert(base string) map[string]string {

	out := map[string]string{}

	for attr := range s {
		out[base+attr] = s[attr]
	}

	return out
}

// DangerouslySetInnerHTMLFunc is a convience function used for setting the DOM
// object's inner html. The functon takes a function for the argument.
//
// See: https://reactjs.org/docs/dom-elements.html#dangerouslysetinnerhtml
func DangerouslySetInnerHTMLFunc(inside func() interface{}) map[string]interface{} {
	return map[string]interface{}{
		"dangerouslySetInnerHTML": map[string]interface{}{
			"__html": inside(),
		},
	}
}

// DangerouslySetInnerHTML is a convience function used for setting the DOM
// object's inner html. The function takes the inner html content directly.
//
// See: https://reactjs.org/docs/dom-elements.html#dangerouslysetinnerhtml
func DangerouslySetInnerHTML(inside interface{}) map[string]interface{} {
	return DangerouslySetInnerHTMLFunc(func() interface{} { return inside })
}

// AddClass adds a new class to an existing list of classes.
func AddClass(currentClasses, newClass string) string {

	uniq := splitClasses(currentClasses)
	uniq[newClass] = struct{}{}

	pre := []string{}
	for k := range uniq {
		pre = append(pre, k)
	}

	return strings.Join(pre, " ")
}

// RemoveClass removes a class from an existing list of classes.
func RemoveClass(currentClasses, removeClass string) string {

	uniq := splitClasses(currentClasses)
	delete(uniq, removeClass)

	pre := []string{}
	for k := range uniq {
		pre = append(pre, k)
	}

	return strings.Join(pre, " ")
}

// splitClasses returns a map where all the keys are the
// unique classes.
func splitClasses(currentClasses string) map[string]struct{} {
	uniq := map[string]struct{}{}

	splits := strings.Split(currentClasses, " ")
	for _, split := range splits {
		split = strings.TrimSpace(split)
		if split != "" {
			uniq[split] = struct{}{}
		}
	}

	return uniq
}
