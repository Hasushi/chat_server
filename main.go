package main

import (
	"fmt"
	"log"
	"net/http"
)

	func main() {
	    fmt.Println("Chat Server Started")

	    // Define HTTP route for the root endpoint
	    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	        fmt.Fprintf(w, "Welcome to the Chat Server")
	    })

	    // TODO: Add websocket handler for chat functionality
	    // http.HandleFunc("/ws", handleConnections)

	    // Start the server
	    log.Println("Server starting on :8080")
	    if err := http.ListenAndServe(":8080", nil); err != nil {
	        log.Fatal("Server error:", err)
	    }

		
	}

