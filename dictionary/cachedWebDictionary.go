package dictionary

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

// CachedWebDictionary uses a REST API to check if
// the given string is a word, caching the result.
type CachedWebDictionary struct {
	cache map[string]bool
}

// IsWord implements Dictionary.IsWord
func (dict *CachedWebDictionary) IsWord(str string) (isWord bool) {
	isWord = false
	if len(str) == 0 {
		return
	}

	var cached bool
	if isWord, cached = dict.cache[str]; cached {
		return
	}

	isWord = isWordInOxfordDictionary(str)
	dict.cache[str] = isWord

	return
}

const (
	endpoint = "lemmas"
	langCode = "en-us"
	// Format with <endpoint>, <langCode>, <word_id>
	oxfordAPIURLFormat = "https://od-api.oxforddictionaries.com/api/v2/%v/%v/%v"
)

var (
	appID  = os.Getenv("OXFORD_APP_ID")
	appKey = os.Getenv("OXFORD_APP_KEY")
)

func isWordInOxfordDictionary(wordID string) bool {
	// Build a request to the Oxford API to see if
	// the word exists in their dictionary.
	url := fmt.Sprintf(oxfordAPIURLFormat, endpoint, langCode, strings.ToLower(wordID))

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("app_id", appID)
	req.Header.Add("app_key", appKey)

	// Send the request.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Handle the response.
	// Status codes defined at:
	// https://developer.oxforddictionaries.com/documentation/response-codes
	switch resp.StatusCode {
	case http.StatusOK:
		return true
	case http.StatusNotFound:
		return false
	default:
		msg := fmt.Sprintf(
			"Received unexpected status code from Oxford API:\n"+
				"\turl: "+url+
				"\tstatus code: %v", resp.StatusCode)
		panic(msg)
	}
}
