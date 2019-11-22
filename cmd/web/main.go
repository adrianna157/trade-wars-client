package main

import (
	"log"
	"net/http"
	"os"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "4000"
	}
	return port
}

func main() {
	testFakeShip()

	port := getPort()
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.HandleFunc("/players", players)
	mux.HandleFunc("/", redirect)
	mux.HandleFunc("/map", showNavigationScreen)
	mux.HandleFunc("/map/trade", showTradeScreen)
	mux.HandleFunc("/map/chat", showChatScreen)
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on " + port)
	err := http.ListenAndServe(":"+port, mux)
	log.Fatal(err)
}
