package dictionary

// CachedWebDictionary uses a REST API to check if
// the given string is a word, caching the result.
type CachedWebDictionary struct {
	cache map[string]bool
}

// IsWord implements Dictionary.IsWord
func (dict *CachedWebDictionary) IsWord(str string) bool {
	return false
}
