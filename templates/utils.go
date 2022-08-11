package templates

import "strings"

// quotes a string if it contains whitespace
func maybeQuote(in string) string {
	if strings.ContainsAny(in, " \n\t") {
		return `"` + in + `"`
	}
	return in
}
