package persian

import "strings"

var (
	plMap = map[rune]rune{
		'0': '۰',
		'1': '۱',
		'2': '۲',
		'3': '۳',
		'4': '۴',
		'5': '۵',
		'6': '۶',
		'7': '۷',
		'8': '۸',
		'9': '۹',
	}

	lpMap = map[rune]rune{
		'۰': '0',
		'۱': '1',
		'۲': '2',
		'۳': '3',
		'۴': '4',
		'۵': '5',
		'۶': '6',
		'۷': '7',
		'۸': '8',
		'۹': '9',
	}

	pDigits = []rune{'۰', '۱', '۲', '۳', '۴', '۵', '۶', '۷', '۸', '۹'}
)

// ToPersianDigits converts all Latin digits in a string to Persian digits.
func ToPersianDigits(str string) string {
	var sb strings.Builder
	for _, r := range []rune(str) {
		if pr, ok := plMap[r]; ok {
			sb.WriteRune(pr)
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
		if ar, ok := lpMap[r]; ok {
			sb.WriteRune(ar)
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

// IsPersianDigit checks if a rune is a Persian digit.
func IsPersianDigit(r rune) bool {
	for _, d := range pDigits {
		if d == r {
			return true
		}
	}
	return false
}
