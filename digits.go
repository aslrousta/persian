package persian

import "strings"

// ToPersianDigits converts all Latin digits in a string to Persian digits.
func ToPersianDigits(str string) string {
	var sb strings.Builder
	for _, r := range []rune(str) {
		if '0' <= r && r <= '9' {
			sb.WriteRune('\u06F0' + r - '0')
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

// ToLatinDigits converts all Persian digits in a string to Latin digits.
func ToLatinDigits(str string) string {
	var sb strings.Builder
	for _, r := range []rune(str) {
		if IsPersianDigit(r) {
			sb.WriteRune('0' + r - '\u06F0')
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

// IsPersianDigit checks if a rune is a Persian digit.
func IsPersianDigit(r rune) bool {
	return '\u06F0' <= r && r <= '\u06F9'
}
