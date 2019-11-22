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

	// // Start the server on localhost port 8000 and log any errors
	// log.Println("http server started on :8000")
	// err := http.ListenAndServe(":8000", nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }

	// // Start the server on localhost port 8000 and log any errors
	// log.Println("http server started on :8000")
	// err := http.ListenAndServe(":8000", nil)
	// if err != nil {
	// 		log.Fatal("ListenAndServe: ", err)
	//   }

	port := getPort()
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.HandleFunc("/players", players)
	mux.HandleFunc("/", redirect)
	mux.HandleFunc("/map", showNavigationScreen)
	mux.HandleFunc("/map/trade", showTradeScreen)
	mux.HandleFunc("/map/chat", showChatScreen)
	// Configure websocket route
	mux.HandleFunc("/map/chatroom", ws)
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Start listening for incoming chat messages
	go handleMessages()

	log.Println("Starting server on " + port)
	err := http.ListenAndServe(":"+port, mux)
	log.Fatal(err)
}
