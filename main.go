package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func fileToByte(fileName string) (bytes []byte) {
	file, _ := os.Open(fileName)
	bytes, _ = ioutil.ReadAll(file)
	return
}

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
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
	w.Write([]byte("This is the snippet"))
}

// Add a createSnippet handler function.
func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("creates new snippit"))
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
