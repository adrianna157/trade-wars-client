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
	// testFakeShip()

	port := getPort()
	// port := os.Getenv("PORT")
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.HandleFunc("/players", players)
	mux.HandleFunc("/", redirect)
	mux.HandleFunc("/map", showMapScreen)
	mux.HandleFunc("/map/trade", showTradeScreen)
	mux.HandleFunc("/map/chat", showChatScreen)
	mux.HandleFunc("/map/moveLeft", moveLeft)
	mux.HandleFunc("/map/moveRight", moveRight)
	mux.HandleFunc("/map/moveUp", moveUp)
	mux.HandleFunc("/map/moveDown", moveDown)
	// Configure websocket route
	fs := http.FileServer(http.Dir("../pkg"))
	http.Handle("/", fs)
	http.HandleFunc("/map/chatroom/ws", handleConnections)
	mux.HandleFunc("/map/chatroom", ws)
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Start listening for incoming chat messages
	go handleMessages()

	log.Println("Starting server on " + port)
	err := http.ListenAndServe(":"+port, mux)
	log.Fatal(err)
}
