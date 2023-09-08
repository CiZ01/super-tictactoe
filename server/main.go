package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ws", handleWebSocketConnection)

	// Avvia il server WebSocket sulla porta 8080
	fmt.Println("Server WebSocket in ascolto su :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
