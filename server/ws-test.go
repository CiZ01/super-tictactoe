package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// ONLY FOR TESTING PURPOSES | print lobbies
func handlePrintLobbies(lobbies []Lobby) {
	fmt.Printf("Lobbies: %+v", lobbies)
}

// ONLY FOR TESTING PURPOSES | print connections
func handlePrintConnections(connections map[*websocket.Conn]connectionStatus) {
	fmt.Printf("Connections: %+v", connections)
}
