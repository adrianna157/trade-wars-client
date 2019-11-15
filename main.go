package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome Screen!"))
}

func showNavigationScreen(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Navigation Screen"))
}

// Add a showSnippet handler function.
func showTradeScreen(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Trade Screen!"))
}

// Add a createSnippet handler function.
func showChatScreen(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Chat Screen!"))
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/navigation", showNavigationScreen)
    mux.HandleFunc("/navigation/trade", showTradeScreen)
    mux.HandleFunc("/naivigation/trade/chat", showChatScreen)


	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil.
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}