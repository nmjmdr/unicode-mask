package masker

import (
	"strings"
	"unicode/utf8"
)

// Mask masks a given string with a given mask character
func Mask(str string, skip int, maskChar rune) string {
	length := utf8.RuneCountInString(str)
	if skip <= 0 || skip >= length {
		return str
	}
	clearString := string([]rune(str)[0:skip])
	return PadRight(clearString, maskChar, length-skip)
}

// PadLeft pads a string on left given string with given char
func PadLeft(str string, r rune, n int) string {
	pad := strings.Repeat(string(r), n)
	return pad + str
}

// PadRight pads a string on right given string with given char
func PadRight(str string, r rune, n int) string {
	pad := strings.Repeat(string(r), n)
	return str + pad
}

// MaskReverse masks a given string with a given mask character, but in reverse order
func MaskReverse(str string, skip int, maskChar rune) string {
	length := utf8.RuneCountInString(str)
	if skip <= 0 || skip >= length {
		return str
	}
	clearString := string([]rune(str)[length-skip:])
	return PadLeft(clearString, maskChar, length-skip)
}

