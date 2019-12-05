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

// func handleConnections(w http.ResponseWriter, r *http.Request) {
// 	// Upgrade initial GET request to a websocket
// 	ws, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// Make sure we close the connection when the function returns
// 	defer ws.Close()

// 	// Register our new client
// 	clients[ws] = true

// 	for {
// 		var msg Message
// 		// Read in a new message as JSON and map it to a Message object
// 		err := ws.ReadJSON(&msg)
// 		if err != nil {
// 			log.Printf("error: %v", err)
// 			delete(clients, ws)
// 			break
// 		}
// 		// Send the newly received message to the broadcast channel
// 		broadcast <- msg
// 	}
// }

// func handleMessages() {
// 	for {
// 		// Grab the next message from the broadcast channel
// 		msg := <-broadcast
// 		// Send it out to every client that is currently connected
// 		for client := range clients {
// 			err := client.WriteJSON(msg)
// 			if err != nil {
// 				log.Printf("error: %v", err)
// 				client.Close()
// 				delete(clients, client)
// 			}
// 		}
// 	}
// }

func main() {
	// testFakeShip()
	go handleMessages()
	

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
	mux.HandleFunc("/addCargo", addCargo)
	// Configure websocket route
	fs := http.FileServer(http.Dir("../pkg"))
	http.Handle("/", fs)
	http.HandleFunc("/map/chatroom/ws", handleConnections)
	mux.HandleFunc("/map/chatroom", ws)
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))


	log.Println("Starting server on " + port)
	err := http.ListenAndServe(":"+port, mux)
	log.Fatal(err)
}
