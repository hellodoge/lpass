package lib

import "strings"

const (
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	numbers = "0123456789"
	special = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
)

func GetChars(uA, ua, u0, u_ bool) string {
	var chars strings.Builder

	if !uA {
		chars.WriteString(uppercase)
	}
	if !ua {
		chars.WriteString(lowercase)
	}
	if !u0 {
		chars.WriteString(numbers)
	}
	if !u_ {
		chars.WriteString(special)
	}

	return chars.String()
}
