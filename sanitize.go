package persian

import "strings"

const (
	percentSign  = '\u066A'
	questionSign = '\u061F'
)

// Sanitizer defines a rune sanitizer.
type Sanitizer interface {
	Sanitize(r rune) rune
}

// Sanitizers
var (
	// Arabic replaces Arabic digits with Persian digits.
	Arabic Sanitizer

	// Special replaces special symbols with Persian counterparts.
	Special Sanitizer
)

func init() {
	Arabic = arabicSanitizer{}
	Special = specialSanitizer{}
}

// Sanitize cleans a string from unwanted characters and symbols.
func Sanitize(str string, sanitizers ...Sanitizer) string {
	var sb strings.Builder
	for _, r := range str {
		for _, s := range sanitizers {
			r = s.Sanitize(r)
		}
		sb.WriteRune(r)
	}
	return sb.String()
}

// IsArabicDigit checks if a rune is an Arabic digit.
func IsArabicDigit(r rune) bool {
	return '\u0660' <= r && r <= '\u0669'
}

type arabicSanitizer struct{}

func (arabicSanitizer) Sanitize(r rune) rune {
	if IsArabicDigit(r) {
		return '\u06F0' + r - '\u0660'
	}
	switch r {
	case '\u064A', '\u0649':
		return '\u06CC'
	case '\u0643':
		return '\u06A9'
	}
	return r
}

type specialSanitizer struct{}

func (specialSanitizer) Sanitize(r rune) rune {
	switch r {
	case '%':
		return percentSign
	case '?':
		return questionSign
	default:
		return r
	}
}
