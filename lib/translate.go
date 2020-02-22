package lib

import "strings"

func Translate(source []byte, chars string) string {
	var out strings.Builder

	for _, i := range source {
		out.WriteByte(chars[int(i) % len(chars)])
	}

	return out.String()
}
