package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/websocket"
)
var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message) 


// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}


type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "4000"
	}
	return port
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	clients[ws] = true

	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}


func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {



	
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
