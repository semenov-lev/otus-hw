package hw02unpackstring

import (
	"errors"
	"strings"
)

var (
	ErrInvalidString = errors.New("invalid string")
	ErrInvalidEscape = errors.New("invalid escape")
)

func Unpack(rawStr string) (result string, err error) {
	stringRunes := []rune(rawStr)
	var strBuilder strings.Builder
	strBuilder.Grow(len(stringRunes))

	for i := 0; i < len(stringRunes); i++ {
		currentRune := stringRunes[i]

		switch {
		case isDigit(currentRune):
			return "", ErrInvalidString
		case currentRune == '\\':
			if i+1 >= len(stringRunes) || !isEscapable(stringRunes[i+1]) {
				return "", ErrInvalidEscape
			}
			i++
			currentRune = stringRunes[i]
		}

		repeatCount := 1
		if i+1 < len(stringRunes) && isDigit(stringRunes[i+1]) {
			if i+2 < len(stringRunes) && isDigit(stringRunes[i+2]) {
				return "", ErrInvalidString
			}
			repeatCount = int(stringRunes[i+1] - '0')
			i++
		}
		strBuilder.WriteString(strings.Repeat(string(currentRune), repeatCount))
	}
	return strBuilder.String(), nil
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isEscapable(r rune) bool {
	return isDigit(r) || isEscape(r)
}

func isEscape(r rune) bool {
	return r == '\\'
}
