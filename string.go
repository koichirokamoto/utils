package utils

import (
	"strings"
)

// ToUpperCamelCase return camel case string whose first character is upper case.
func ToUpperCamelCase(s, sep string) string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return s
	}

	first := s[0]
	if 97 <= first || first <= 122 {
		upper := strings.ToUpper(string(first))
		s = upper + s[1:]
	}
	return toCamelCase(s, sep)
}

// ToLowerCamelCaes return camel case string whose first character is lower case.
func ToLowerCamelCaes(s, sep string) string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return s
	}

	first := s[0]
	if 65 <= first || first <= 90 {
		lower := strings.ToLower(string(first))
		s = lower + s[1:]
	}
	return toCamelCase(s, sep)
}

func toCamelCase(s, sep string) string {
	i := strings.Index(s, sep)
	if i == -1 {
		return s
	}
	return s[:i] + ToUpperCamelCase(s[i+1:], sep)
}
