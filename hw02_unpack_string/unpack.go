package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")
var ErrInvalidEscape = errors.New("invalid escape")

func Unpack(rawStr string) (result string, err error) {
	stringRunes := []rune(rawStr)
	resultStrings := make([]string, 0, len(stringRunes))
	var escapeCase bool

	for i := 0; i < len(stringRunes); i++ {
		current := stringRunes[i]
		var next rune

		if i+1 < len(stringRunes) {
			next = stringRunes[i+1]
		}

		currentIsDigit := unicode.IsDigit(current)
		nextIsDigit := unicode.IsDigit(next)
		currentIsEscaping := current == '\\'
		nextIsEscaping := next == '\\'
		switch {
		case !escapeCase && unicode.IsDigit(current) && nextIsDigit:
			err = ErrInvalidString
			return
		case escapeCase:
			if current == 'n' && nextIsDigit {
				digit, _ := strconv.Atoi(string(next))
				resultStrings = append(resultStrings, strings.Repeat(string('\\')+string(current), digit))
				escapeCase = false
				i++
			} else if currentIsDigit {
				if nextIsDigit {
					digit, _ := strconv.Atoi(string(next))
					resultStrings = append(resultStrings, strings.Repeat(string(current), digit))
					escapeCase = false
					i++
				} else if nextIsEscaping {
					resultStrings = append(resultStrings, string(current))
					escapeCase = false
				} else {
					resultStrings = append(resultStrings, string(current))
					escapeCase = false
				}
			} else if nextIsDigit {
				if currentIsEscaping {
					digit, _ := strconv.Atoi(string(next))
					resultStrings = append(resultStrings, strings.Repeat(string(current), digit))
				} else {
					err = ErrInvalidString
					return
				}
				i++
			} else if nextIsEscaping {
				resultStrings = append(resultStrings, string('\\'))
				escapeCase = false
			} else {
				err = ErrInvalidEscape
				return
			}
		case currentIsEscaping:
			escapeCase = true
		case currentIsDigit:
			err = ErrInvalidString
			return
		case nextIsDigit:
			digit, _ := strconv.Atoi(string(next))
			resultStrings = append(resultStrings, strings.Repeat(string(current), digit))
			i++
		default:
			resultStrings = append(resultStrings, string(current))
		}
	}

	defer func() {
		result = strings.Join(resultStrings, "")
	}()

	return result, err
}
