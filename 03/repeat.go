package iteration

import "strings"

// Repeat repeats a string N times
func Repeat(character string, N int) string {

	var repeated string
	for i := 0; i < N; i++ {
		repeated += character
	}

	return repeated
}

// RepeatWithStrings repeats a string N times using the 'strings' package
func RepeatWithStrings(character string, N int) string {
	return strings.Repeat(character, N)
}
