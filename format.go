package persian

import (
	"fmt"
	"strings"
)

const (
	decimalSeparator   = '\u066B'
	thousandsSeparator = '\u066C'
)

// Wrap marks a value v to be formatted with Persian locale.
func Wrap(v interface{}) fmt.Stringer {
	return wrapper{v, false}
}

// WrapWithSep marks a value v to be formatted with Persian locale
// and thousands separators between digits.
func WrapWithSep(v interface{}) fmt.Stringer {
	return wrapper{v, true}
}

type wrapper struct {
	v   interface{}
	sep bool
}

func (ws wrapper) String() string {
	str := ToPersianDigits(fmt.Sprintf("%v", ws.v))
	return format(str, ws.sep)
}

func format(str string, sep bool) string {
	const (
		NonDigit = iota
		DigitWithSep
		DigitWithPeriod
		DigitWithoutSep
	)

	state := NonDigit

	var sb strings.Builder
	var digits []rune

	shipDigits := func() {
		for len(digits) % 3 > 0 {
			sb.WriteRune(digits[0])
			digits = digits[1:]
		}
		for len(digits) > 0 {
			if sep {
				sb.WriteRune(thousandsSeparator)
			}
			sb.WriteRune(digits[0])
			sb.WriteRune(digits[1])
			sb.WriteRune(digits[2])
			digits = digits[3:]
		}
	}

	for _, r := range []rune(str) {
		switch state {
		case NonDigit:
			if IsPersianDigit(r) {
				digits = []rune{r}
				state = DigitWithSep
			} else {
				sb.WriteRune(r)
			}
		case DigitWithSep:
			if IsPersianDigit(r) {
				digits = append(digits, r)
			} else if r == '.' {
				shipDigits()
				state = DigitWithPeriod
			} else {
				shipDigits()
				state = NonDigit
			}
		case DigitWithPeriod:
			if IsPersianDigit(r) {
				sb.WriteRune(decimalSeparator)
				sb.WriteRune(r)
				state = DigitWithoutSep
			} else {
				sb.WriteRune('.')
				sb.WriteRune(r)
				state = NonDigit
			}
		case DigitWithoutSep:
			sb.WriteRune(r)
			if !IsPersianDigit(r) {
				state = NonDigit
			}
		}
	}

	shipDigits()

	return sb.String()
}
