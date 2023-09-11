package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func handler(messageType int, p []byte, conn *websocket.Conn) {
	switch messageType {
	case websocket.TextMessage:
		var message map[string]interface{}
		if err := json.Unmarshal(p, &message); err != nil {
			log.Println("Errore nell'analisi del messaggio JSON:", err)
			return
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
		// TEST PURPOSES ONLY
		case "print_lobbies":
			handlePrintLobbies(lobbies)
		case "print_connections":
			handlePrintConnections(connections)
		default:
			log.Println("Tipo di messaggio non valido:", messageType)
		}
	}
}

// just for testing
func handleRegistration(conn *websocket.Conn, message map[string]interface{}) {
	newUserUsername := message["username"].(string)

	// create a new user
	newUser := User{
		Conn:     conn,
		Username: newUserUsername,
	}

	// Registra la connessione
	connections[conn] = connectionStatus{StatusCode: 1, User: newUser}

	// Invia un messaggio di conferma di registrazione al giocatore
	response := map[string]interface{}{
		"type":    "registration_success",
		"message": "Registrazione completata con successo",
	}
	fmt.Printf("Connessione %v registrata con successo\n", newUser)
	conn.WriteJSON(response)
}

func handleNewLobby(conn *websocket.Conn, message map[string]interface{}) {
	// Esegui la logica del gioco qui
	for _, l := range lobbies {
		if u := l.GetUserByConn(conn); u != nil {
			l.AddUser(&User{ID: "2", Conn: conn})
			response := map[string]interface{}{
				"type":    "added_to_lobby_confirmation",
				"message": "Aggiunto alla lobby con successo",
			}
			conn.WriteJSON(response)
			return
		}
	}

	newLobby := NewLobby()
	newLobby.AddUser(&User{ID: "1", Conn: conn})
	lobbies = append(lobbies, *newLobby)

	response := map[string]interface{}{
		"type":    "new_lobby_confirmation",
		"message": "Nuova lobby creata con successo",
	}

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
