package main

import (
	"log"
	"net/http"
	"os"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://localhost:ex4000/players", 301)
}

func main() {
	http.HandleFunc("/", redirect)

	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.

	mux := http.NewServeMux()
	mux.HandleFunc("/players", players)
	mux.HandleFunc("/", redirect)
	mux.HandleFunc("/map", showNavigationScreen)
	mux.HandleFunc("/map/trade", showTradeScreen)
	mux.HandleFunc("/map/chat", showChatScreen)
	mux.HandleFunc("/snippet/create", createSnippet)

	mux.HandleFunc("/snippet", showSnippet)
	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir function is relative to the project
	// directory root.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil.
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), mux)
	log.Fatal(err)
}
