package stringutils

import (
	"fmt"
	"github.com/dbelc/camelCase/dictionary"
	"strings"
	"unicode"
)

// CamelCase applies the camelCase function
func CamelCase(in string, dict dictionary.Dictionary) (string, error) {
	lower := strings.ToLower(in)
	runes := []rune(lower)

	// Sentinel to keep track of first word found or not.
	// The first word requires special casing.
	isFirstWord := true

	// Iterate over the string, checking each index for
	// maximal length word substrings and applying camelCasing
	for wordStart := 0; wordStart < len(runes); wordStart++ {
		wordLength := findLongestWordIn(runes[wordStart:], dict)
		if wordLength == -1 {
			// String contains non-words. Return an error.
			return "", fmt.Errorf(
				"String contains non-words, starting with character %v, substring: %s",
				wordStart+1, string(runes[wordStart:]))
		}

		if isFirstWord {
			// String is already lowercase so we continue
			isFirstWord = false
		} else {
			// Word should be Uppercase
			runes[wordStart] = unicode.ToUpper(runes[wordStart])
		}

		// Set wordStart to after the identified word.
		// Set to 1 less than the correct value because
		// the loop will also increment wordStart by 1.
		wordStart += wordLength - 1
	}

	return string(runes), nil
}

func findLongestWordIn(runes []rune, dict dictionary.Dictionary) int {
	runesLen := len(runes)
	for length := runesLen; length >= 0; length-- {
		substr := string(runes[0:length])
		if dict.IsWord(substr) {
			return length
		}
	}

	return -1
}
