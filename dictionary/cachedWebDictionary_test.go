package dictionary

import "testing"

type subtestFn = func(*testing.T, *CachedWebDictionary)

type subtest struct {
	name   string
	testFn subtestFn
}

func TestCachedWebDictionary(t *testing.T) {
	for _, subtest := range []subtest{
		{"Valid words", validWords},
		{"Cached valid words", cachedValidWords},
		{"Invalid words", invalidWords},
	} {
		t.Run(subtest.name, func(*testing.T) {
			subtest.testFn(t, newCachedWebDictionary())
		})
	}
}

// Test helpers

func testIsWord(want bool, t *testing.T, dict *CachedWebDictionary, cases []string) {
	for _, c := range cases {
		isWord := dict.IsWord(c)
		t.Errorf("CachedWebDictionary.IsWord(%q) == %v, want %v", c, isWord, want)
	}
}

// Test cases

func validWords(t *testing.T, dict *CachedWebDictionary) {
	testIsWord(true, t, dict, []string{
		// Basic test for valid dictionary words
		"one",
		"two",
		"red",
		"blue",
		"couch",
		"keyboard",
	})
}

func cachedValidWords(t *testing.T, dict *CachedWebDictionary) {
	// Ensure the cache is empty
	dict.cache = make(map[string]bool)

	// Test the cache
	validWordsList := []string{
		"candle",
		"controller",
		"sink",
		"book",
		"xylophone",
	}

	for _, word := range validWordsList {
		var isWord bool
		if isWord = dict.IsWord(word); !isWord {
			t.Errorf("CachedWebDictionary.IsWord(%q) == false, want true", word)
		} else {
			// Test that the word was cached
			cachedIsWord, prs := dict.cache[word]
			if !prs {
				t.Errorf("CachedWebDictionary.cache expected to contain key %q but did not", word)
			} else if isWord != cachedIsWord {
				t.Errorf("CachedWebDictionary.cache[%q] == %v, want %v", word, cachedIsWord, isWord)
			}
		}
	}

	invalidWordsList := []string{
		"thisisasentence",
		"awefawef",
		"1123198",
		"",
	}

	for _, word := range invalidWordsList {
		var isWord bool
		if isWord = dict.IsWord(word); isWord {
			t.Errorf("CachedWebDictionary.IsWord(%q) == true, want false", word)
		} else {
			// Test the word was not cached
			if _, prs := dict.cache[word]; prs {
				t.Errorf("CachedWebDictionary.cache expected to not contain key %q but did ", word)
			}
		}
	}
}

func invalidWords(t *testing.T, dict *CachedWebDictionary) {
	testIsWord(false, t, dict, []string{
		"",
		"thisismanywords",
		"fawoefjasawef",
		"112312312",
		"almost1",
		"closeawefaw",
	})
}
