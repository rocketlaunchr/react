package main

import (
	"strings"
)

// Set is used for conveniently dealign with
// data-* and aria-* attributes.
// See: https://reactjs.org/docs/dom-elements.html
type Set map[string]string

// Convert converts a Set into a map containing the actual
// attribute names by prefixing the base.
func (s Set) Convert(base string) map[string]string {

	out := map[string]string{}

	for attr := range s {
		out[base+attr] = s[attr]
	}

	return out
}

// EscapeHTMLFunc is used to create a prop that can escape and set
// the inner html of an element.
func EscapeHTMLFunc(inside func() interface{}) map[string]interface{} {
	return map[string]interface{}{
		"dangerouslySetInnerHTML": map[string]interface{}{
			"__html": inside(),
		},
	}
}

// EscapeHTML is a convience function.
func EscapeHTML(inside interface{}) map[string]interface{} {
	return EscapeHTMLFunc(func() interface{} { return inside })
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
func RemoveClass(currentClasses, newClass string) string {

	uniq := splitClasses(currentClasses)
	delete(uniq, newClass)

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
