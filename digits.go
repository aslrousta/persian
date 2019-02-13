package persian

import "strings"

var (
	paMap = map[rune]rune{
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

	apMap = map[rune]rune{
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
)

// ToPersianDigits converts all ASCII digits in a string to Persian
// digits.
func ToPersianDigits(str string) string {
	sb := new(strings.Builder)
	for _, r := range []rune(str) {
		if pr, ok := paMap[r]; ok {
			sb.WriteRune(pr)
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

// ToASCIIDigits converts all Persian digits in a string to ASCII
// digits.
func ToASCIIDigits(str string) string {
	sb := new(strings.Builder)
	for _, r := range []rune(str) {
		if ar, ok := apMap[r]; ok {
			sb.WriteRune(ar)
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}
