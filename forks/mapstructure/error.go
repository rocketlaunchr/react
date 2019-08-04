package mapstructure

import (
	"errors"
	"sort"
	"strings"

	"github.com/rocketlaunchr/react/forks/fmtless"
)

func sprintf(fmts string, args ...interface{}) (rStr string) {
	defer func() {
		if r := recover(); r != nil {
			// rStr = r.(string)
			println("fmtless error: " + r.(string))
		}
	}()
	return fmt.Sprintf(fmts, args...)
}

func errorf(format string, a ...interface{}) (rErr error) {
	defer func() {
		if r := recover(); r != nil {
			println("fmtless error: " + r.(string))
			rErr = errors.New(r.(string))
		}
	}()
	return fmt.Errorf(format, a...)
}

// Error implements the error interface and can represents multiple
// errors that occur in the course of a single decode.
type Error struct {
	Errors []string
}

func (e *Error) Error() string {
	points := make([]string, len(e.Errors))
	for i, err := range e.Errors {
		points[i] = sprintf("* %s", err)
	}

	sort.Strings(points)
	return sprintf(
		"%d error(s) decoding:\n\n%s",
		len(e.Errors), strings.Join(points, "\n"))
}

// WrappedErrors implements the errwrap.Wrapper interface to make this
// return value more useful with the errwrap and go-multierror libraries.
func (e *Error) WrappedErrors() []error {
	if e == nil {
		return nil
	}

	result := make([]error, len(e.Errors))
	for i, e := range e.Errors {
		result[i] = errors.New(e)
	}

	return result
}

func appendErrors(errors []string, err error) []string {
	switch e := err.(type) {
	case *Error:
		return append(errors, e.Errors...)
	default:
		return append(errors, e.Error())
	}
}
