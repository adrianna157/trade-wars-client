package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func fileToByte(fileName string) (bytes []byte) {
	file, _ := os.Open(fileName)
	bytes, _ = ioutil.ReadAll(file)
	return
}

func players(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fileToByte("welcome-screen.html")))
}

func showNavigationScreen(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fileToByte("navigation-screen.html")))
}

// Add a showSnippet handler function.
func showTradeScreen(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fileToByte("trade-screen.html")))
}

// Add a createSnippet handler function.
func showChatScreen(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fileToByte("chat-screen.html")))
}

// Add a showSnippet handler function.
func showSnippet(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string and try to
	// convert it to an integer using the strconv.Atoi() function. If it can't
	// be converted to an integer, or the value is less than 1, we return a 404 page
	// not found response.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// Use the fmt.Fprintf() function to interpolate the id value with our response
	// and write it to the http.ResponseWriter.
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Add a createSnippet handler function.
func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// If it's not, use the w.WriteHeader() method to send a 405 status
		// code and the w.Write() method to write a "Method Not Allowed"
		// response body. We then return from the function so that the
		// subsequent code is not executed.
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}
