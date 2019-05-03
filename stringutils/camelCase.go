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

	split, err := splitIntoWords(runes, dict)
	if err != nil {
		return "", err
	}

	// Iterate through the list of splits and camelCase the words
	// Keep a running sum of wordLen
	wordLenSum := 0
	for i, wordLen := range split {
		if i == 0 {
			// String is already lowercase, nothing to do
		} else {
			runes[wordLenSum] = unicode.ToUpper(runes[wordLenSum])
		}

		wordLenSum += wordLen
	}

	return string(runes), nil
}

// Splits a slice of runes into words.
// Returns a slice of ints where each value is the length of
//   the next valid word.
func splitIntoWords(runes []rune, dict dictionary.Dictionary) ([]int, error) {
	if len(runes) == 0 {
		return make([]int, 0), nil
	}

	for i := range runes {
		wordLen := i + 1
		substr := string(runes[0:wordLen])
		if dict.IsWord(substr) {
			if split, err := splitIntoWords(runes[wordLen:], dict); err == nil {
				// split is valid, add our wordLen to the front
				newSplit := make([]int, 1)
				newSplit[0] = wordLen

				newSplit = append(newSplit, split...)
				return newSplit, nil
			}
		}
	}

	return make([]int, 0), fmt.Errorf("No valid split was possible")
}
