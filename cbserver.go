package main

import (
	"fmt"
	"net/http"
)

// A simple http callback server for spotify auth
func main() {

	fmt.Println("Starting callback server on port 8731")
	fmt.Println("Listening on http://localhost:8731/callback")
	fmt.Println("Press Ctrl+C to exit")

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("callback received", r.URL.Query()["code"])
		fmt.Fprintf(w, r.URL.Query()["code"][0])
	})

	http.ListenAndServe(":8731", nil)
}
