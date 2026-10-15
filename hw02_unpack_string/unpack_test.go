package hw02unpackstring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: "🙃0", expected: ""},
		{input: "aaф0b", expected: "aab"},
		{input: "d\n5abc", expected: "d\n\n\n\n\nabc"},
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
		{input: "🙃3я2", expected: "🙃🙃🙃яя"},
		{input: "世2界", expected: "世世界"},
		{input: `\\`, expected: `\`},
		{input: `\40`, expected: ``},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{
		"3abc",
		"45",
		"aaa10b",
		`\455`,
		`Hello, 世界\\33`,
		`Hello, 世界\\\\33`,
	}
	for _, tc := range invalidStrings {
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.ErrorIs(t, err, ErrInvalidString)
		})
	}
}

func TestUnpackInvalidEscape(t *testing.T) {
	invalidEscapes := []string{
		`\n`,
		`abc\a`,
		`a\abc`,
		`qw\ne`,
		`\a5`,
		`\a\`,
		`\`,
		`Hello, 世界\`,
		"d\\n5abc",
	}
	for _, tc := range invalidEscapes {
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.ErrorIs(t, err, ErrInvalidEscape)
		})
	}
}
