package dictionary

import "reflect"

// A Dictionary tests if a given string is a word.
type Dictionary interface {
	IsWord(string) bool
}

// New creates a new instance of the requested dictionary type, if possible.
func New(t interface{}) Dictionary {
	name := reflect.ValueOf(t).Type().Name()
	switch name {
	case "CachedWebDictionary":
		return newCachedWebDictionary()
	default:
		panic("Unexpected type: " + name)
	}
}

func newCachedWebDictionary() *CachedWebDictionary {
	return &CachedWebDictionary{make(map[string]bool)}
}
