package main

import "github.com/gorilla/websocket"

// User rappresenta un utente nel gioco.
type User struct {
	ID   string
	Conn *websocket.Conn // Connessione WebSocket dell'utente
}


