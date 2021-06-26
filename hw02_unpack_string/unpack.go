package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var buf strings.Builder
	var prevChar rune

	for i, char := range s {
		if unicode.IsDigit(char) {
			if i == 0 || (unicode.IsDigit(prevChar)) {
				return "", ErrInvalidString
			}
			if unicode.IsPunct(prevChar) {
				buf.WriteRune(char)
				prevChar = char
				continue
			}
			n, _ := strconv.Atoi(string(char))
			repeat := strings.Repeat(string(prevChar), n)
			buf.WriteString(repeat)
		}
		if unicode.IsLetter(char) {
			if unicode.IsLetter(prevChar) {
				buf.WriteRune(prevChar)
			}
			if i == len(s)-1 {
				buf.WriteRune(char)
			}
		}
		if unicode.IsPunct(char) {
			if unicode.IsLetter(prevChar) {
				buf.WriteRune(prevChar)
			}
		}
		prevChar = char
	}

	return buf.String(), nil
}
