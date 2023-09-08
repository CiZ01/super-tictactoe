package main

import (
	"encoding/json"
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
	lobbies = make(map[string]*Lobby)
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

		// Decodifica il messaggio in base al tipo di messaggio
		switch messageType {
		case websocket.TextMessage:
			var message map[string]interface{}
			if err := json.Unmarshal(p, &message); err != nil {
				log.Println("Errore nell'analisi del messaggio JSON:", err)
				continue
			}

			// Gestisci il messaggio in base al suo tipo
			messageType := message["type"].(string)
			switch messageType {
			case "registration":
				handleRegistration(conn, message)
			case "move":
				handleMove(conn, message)
			case "new_lobby":
				handleNewLobby(conn, message)
			case "win":
				handleWin(conn, message)
			default:
				log.Println("Tipo di messaggio non valido:", messageType)
			}
		}

		// Rimuovi la connessione dalla mappa quando il client si disconnette
		connMutex.Lock()
		delete(connections, conn)
		connMutex.Unlock()
	}
}

// just for testing
func handleRegistration(conn *websocket.Conn, message map[string]interface{}) {
	connections[conn] = connectionStatus{StatusCode: 1, User: User{ID: "1", Conn: conn}}

	// Invia un messaggio di conferma di registrazione al giocatore
	response := map[string]interface{}{
		"type":    "registration_success",
		"message": "Registrazione completata con successo",
	}
	conn.WriteJSON(response)
}

func handleNewLobby(conn *websocket.Conn, message map[string]interface{}) {
	// Esegui la logica del gioco qui
	for l := range lobbies{
		if l.
		}
	}
	newLobby := NewLobby()
	newLobby.AddUser(&User{ID: "1", Conn: conn})
	conn.WriteJSON(response)

	// Puoi anche aggiornare gli altri giocatori con lo stato aggiornato del gioco
	// Utilizzando la funzione UpdateGameBoard o UpdateGameStatus che abbiamo discusso prima
}

func handleMove(conn *websocket.Conn, message map[string]interface{}) {
	// Esegui la logica del gioco qui

	// Ad esempio, puoi inviare una risposta al giocatore che ha effettuato la mossa
	response := map[string]interface{}{
		"type": "move_confirmation",
	}
	conn.WriteJSON(response)

	// Puoi anche aggiornare gli altri giocatori con lo stato aggiornato del gioco
	// Utilizzando la funzione UpdateGameBoard o UpdateGameStatus che abbiamo discusso prima
}

func handleWin(conn *websocket.Conn, message map[string]interface{}) {
	// Esegui la logica del gioco qui

	// Ad esempio, puoi inviare una risposta al giocatore che ha effettuato la mossa
	response := map[string]interface{}{
		"type": "win_confirmation",
	}
	conn.WriteJSON(response)

	// Puoi anche aggiornare gli altri giocatori con lo stato aggiornato del gioco
	// Utilizzando la funzione UpdateGameBoard o UpdateGameStatus che abbiamo discusso prima
}
