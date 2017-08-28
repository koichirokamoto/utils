package utils

import (
	"strings"
	"unicode/utf8"
)

// ToUpperCamelCase return camel case string whose first character is upper case.
func ToUpperCamelCase(s, sep string) string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return s
	}

	first := s[0]
	// Check first is full rune because of discarding multi byte character.
	if utf8.FullRune([]byte{first}) && 97 <= first && first <= 122 {
		upper := strings.ToUpper(string(first))
		s = upper + s[1:]
	}
	return toCamelCase(s, sep)
}

// ToLowerCamelCase return camel case string whose first character is lower case.
func ToLowerCamelCase(s, sep string) string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return s
	}

	first := s[0]
	// Check first is full rune because of discarding multi byte character.
	if utf8.FullRune([]byte{first}) && 65 <= first && first <= 90 {
		lower := strings.ToLower(string(first))
		s = lower + s[1:]
	}
	return toCamelCase(s, sep)
}

func toCamelCase(s, sep string) string {
	// If separator is empty, return s.
	if sep == "" {
		return s
	}
	i := strings.Index(s, sep)
	if i == -1 {
		return s
	}
	return s[:i] + ToUpperCamelCase(s[i+1:], sep)
}

// SeparateCamelCase separete camel case string by separator.
func SeparateCamelCase(s string, sep string) string {
	if sep == "" {
		return s
	}

	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	var ns string
	for _, r := range s {
		b := make([]byte, utf8.RuneLen(r))
		utf8.EncodeRune(b, r)
		first := b[0]
		if utf8.FullRune([]byte{b[0]}) && 65 <= first && first <= 90 {
			ns += sep + strings.ToLower(string(r))
			continue
		}
		ns += string(r)
	}

	return ns
}
