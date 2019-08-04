package fmt

import "strings"

// SRepeat just repeats a character or string times <repeat>.
func SRepeat(char string, repeat int) string {
	out := []string{}
	for i := 0; i < repeat; i++ {
		out = append(out, char)
	}
	return strings.Join(out, "")
}
