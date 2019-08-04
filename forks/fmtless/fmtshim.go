package fmt

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

// Errorf returnsan error object for the given format and values.
func Errorf(format string, a ...interface{}) error {
	return errors.New(Sprintf(format, a...))
}

// Printf prints a format string filled with the given values.
func Printf(format string, a ...interface{}) (n int, err error) {
	out := Sprintf(format, a...)
	print(out)
	return len([]byte(out)), nil
}

// Println formats using the default formats for its
// operands and writes to standard output. Spaces are
// always added between operands and a newline is appended.
// It returns the number of bytes written and any write
// error encountered.
func Println(a ...interface{}) (n int, err error) {
	out := Sprint(a...)
	println(out)
	return len([]byte(out)), nil
}

// Print formats using the default formats for its
// operands and writes to standard output. Spaces
// are added between operands when neither is a string.
// It returns the number of bytes written and any
// write error encountered.
func Print(a ...interface{}) (n int, err error) {
	out := Sprint(a...)
	print(out)
	return len([]byte(out)), nil
}

// Sprint renders arguments in their default format
// ("%s"/"%d"/"%f"/"%U" for string/int/float/rune, respectively)
func Sprint(a ...interface{}) string {
	var out []string
	for _, i := range a {
		out = append(out, fmtI("%s", i))
	}
	return strings.Join(out, " ")
}

// Sprintf is a fmtless alternative to fmt.Sprintf that supports some of
// the most common subset of fmt usage.
func Sprintf(fmts string, args ...interface{}) string {
	var bits []string
	fmlist := splitFmtSpecs(fmts)
	for idx, sm := range fmlist {
		var i interface{}
		i = nil
		if idx < len(args) {
			i = args[idx]
		}
		bits = append(bits, sm.render(i))
	}
	return strings.Join(bits, "")
}

// Sprintln is just like Sprint but with a trailing newline.
func Sprintln(a ...interface{}) string {
	return Sprint(a...) + "\n"
}

type sprintMatch struct {
	before string
	spec   string
}

func (spm sprintMatch) render(i ...interface{}) string {
	if spm.spec == "" {
		return spm.before
	}
	return spm.before + fmtI(spm.spec, i[0])
}

func splitFmtSpecs(fmts string) []sprintMatch {
	var (
		lastMatch int
		window    []byte
		matches   []sprintMatch
	)
	percent := byte('%')
	// strings.Index uses a search algo that's
	// probably slower than straight iteration for
	// the length of most format strings.
	for i := 0; i < len(fmts); i++ {
		bound := i + 3
		if bound >= len(fmts) {
			bound = len(fmts)
		}
		window = []byte(fmts)[i:bound]
		if len(window) < 2 || window[0] != percent {
			continue
		}
		spec, ok := getSpec(window)
		if !ok {
			continue
		}
		sm := sprintMatch{before: fmts[lastMatch:i], spec: spec}
		matches = append(matches, sm)
		lastMatch = i + len(sm.spec)
		i += len(sm.spec) - 1
	}
	if lastMatch < len(fmts) {
		lm := sprintMatch{before: fmts[lastMatch:], spec: ""}
		matches = append(matches, lm)
	}
	return matches
}

func getSpec(window []byte) (string, bool) {
	if window[0] != '%' {
		return "", false
	}
	var speclen int
	if window[1] == '+' || window[1] == '#' {
		speclen = 3
	} else {
		speclen = 2
	}
	switch window[speclen-1] {
	case 'v', 's', 'q', 'd', 'b', 'f', 'F', 'g', 'G', 'e', 'E', 'o', 'x', 'X', 'U':
		return string(window[:speclen]), true
	default:
		return "", false
	}
}

func fmtI(spec string, i interface{}) string {
	switch i.(type) {
	case reflect.Type:
		return fmtString(spec, i.(reflect.Type).String())
	case error:
		return fmtString(spec, i.(error).Error())
	case string:
		return fmtString(spec, i.(string))
	case []byte:
		return fmtBytes(spec, i.([]byte))
	case rune:
		return fmtUEscape(i.(rune))
	case int, int64:
		return fmtInt(spec, i)
	case float32, float64:
		return fmtFloat(spec, i)
	case reflect.Value:
		return fmtI(spec, (i.(reflect.Value)).Interface())
	case reflect.Kind:
		panic("unsupported interface for fmtless.Sprintf (reflect.Kind): " + (i.(reflect.Kind)).String())
	default:
		panic("unsupported interface for fmtless.Sprintf: " + reflect.TypeOf(i).String())
	}
}

func fmtBytes(spec string, i []byte) string {
	switch spec {
	case "%v", "%q", "%s":
		return fmtString(spec, string(i))
	case "%x", "%X":
		{
			var b16bytes []string
			for _, b := range i {
				bs := strconv.FormatInt(int64(b), 16)
				if len(bs) == 1 {
					bs = "0" + bs
				}
				b16bytes = append(b16bytes, bs)
			}
			if spec == "%X" {
				return strings.ToUpper(strings.Join(b16bytes, ""))
			}
			return strings.Join(b16bytes, "")
		}
	default:
		panic("Unsupported spec for []byte: " + spec)
	}
}

func fmtString(spec, i string) string {
	switch spec {
	case "%s", "%v", "%#v":
		return i
	case "%q":
		return strconv.Quote(i)
	default:
		panic("Unsupported spec for string: " + spec)
	}
}

func fmtUEscape(r rune) string {
	rcode := strings.ToUpper(strconv.FormatInt(int64(r), 16))
	rcode = "U+" + SRepeat("0", 4-len(rcode)) + rcode
	return rcode
}

func fmtInt(spec string, i interface{}) string {
	var i64 int64
	var base int
	switch i.(type) {
	case int:
		i64 = int64(i.(int))
	case int32:
		i64 = int64(i.(int32))
	case int64:
		i64 = i.(int64)
	}
	switch spec {
	case "%s", "%d", "%v":
		base = 10
	case "%o":
		base = 8
	case "%b":
		base = 2
	case "%X", "%x":
		base = 16
	}
	out := strconv.FormatInt(i64, base)
	if spec == "%X" {
		out = strings.ToUpper(out)
	}
	return out
}

func fmtFloat(spec string, i interface{}) string {
	var (
		bs  int
		ifl float64
	)
	switch i.(type) {
	case float32:
		{
			ifl = float64(i.(float32))
			bs = 32
		}
	case float64:
		{
			ifl = i.(float64)
			bs = 64
		}
	default:
		panic("Unpossible")
	}
	switch spec {
	case "%b", "%f", "%F", "%g", "%G", "%e", "%E":
		return strconv.FormatFloat(ifl, spec[1], -1, bs)
	case "%s", "%v":
		return strconv.FormatFloat(ifl, 'f', -1, bs)
	default:
		panic("Unsupported specifier for floats: " + spec)
	}
}
