package main

import "github.com/gorilla/websocket"

// User rappresenta un utente nel gioco.
type User struct {
	Conn     *websocket.Conn // Connessione WebSocket dell'utente
	Username string          // Username dell'utente
}
