package main

import (
	"fmt"
	"github.com/dbelc/camelCase/dictionary"
	"github.com/dbelc/camelCase/stringutils"
	"log"
	"net/http"
	"os"
)

const (
	defaultPort = "8080"

	camelCasePath         = "/api/camelCase/"
	camelCasePathLen      = len(camelCasePath)
	errorEmptyStringInput = "Invalid input: String cannot be empty."
)

var dict = dictionary.New(dictionary.CachedWebDictionary{})

func camelCaseHandler(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Path[camelCasePathLen:]
	if len(input) == 0 {
		http.Error(w, errorEmptyStringInput, http.StatusBadRequest)
	}

	defer func() {
		if r := recover(); r != nil {
			http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		}
	}()

	camelCase, err := stringutils.CamelCase(input, dict)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Fprintf(w, camelCase)
}

func main() {
	port := defaultPort
	if platformPort := os.Getenv("HTTP_PLATFORM_PORT"); platformPort != "" {
		port = platformPort
	}

	http.HandleFunc(camelCasePath, camelCaseHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
