package main

import (
	"fmt"
	"github.com/dbelc/camelCase/stringutils"
	"net/http"
	"os"
)

func camelCaseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, stringutils.CamelCase("Test"))
}

func main() {
	port := "3001"
	if os.Getenv("HTTP_PLATFORM_PORT") != "" {
		port = os.Getenv("HTTP_PLATFORM_PORT")
	}

	http.HandleFunc("/camelCase/", camelCaseHandler)
	http.ListenAndServe(":"+port, nil)
}
