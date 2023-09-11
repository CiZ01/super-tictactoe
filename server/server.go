package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type connectionStatus struct {
	StatusCode int // 0: not registerd, 1: registered, 2: in lobby
	User       User
}

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true // Per consentire tutte le origini (cambia per ambienti di produzione)
		},
	}
	connections = make(map[*websocket.Conn]connectionStatus)
	connMutex   = sync.Mutex{}

	// Lobby del gioco
	lobbies = []Lobby{}
)

func handleWebSocketConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// Aggiungi la connessione alla mappa delle connessioni
	connMutex.Lock()
	connections[conn] = connectionStatus{StatusCode: 0}
	connMutex.Unlock()

	for {
		// Leggi il messaggio inviato dal client
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		// Gestisci il messaggio in base al suo tipo
		handler(messageType, p, conn)
	}
	// Rimuovi la connessione dalla mappa quando il client si disconnette
	connMutex.Lock()
	delete(connections, conn)
	connMutex.Unlock()
}
