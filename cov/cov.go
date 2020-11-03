package cov

import "strings"

func Words(s string) string {
	c := len(strings.Fields(s))
	switch {
	case c == 0:
		return "0"
	case c == 1:
		return "1"
	case c < 4:
		return "4以下"
	case c < 8:
		return "8ika"
	default:
		return "too many words"
	}
}
