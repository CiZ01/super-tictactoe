package main

import (
	"encoding/json"
	"time"

	"github.com/gorilla/websocket"
)

type GameStatus struct {
	StatusCode int       // 0: in attesa, 1: in corso, 2: terminata
	StartTime  time.Time // Data e ora di inizio della partita
	EndTime    time.Time // Data e ora di fine della partita
}

type Move struct {
	MasterBoardCell          int
	SubBoardCell             int
	MasterBoardCellCompleted bool
}

type NextMove struct {
	Move   Move
	Player string
}

// UpdatenextMove aggiorna la griglia di gioco e la invia a entrambi i giocatori.
func UpdatenextMove(player1, player2 *websocket.Conn, nextMove NextMove) error {
	// Converti nextMove in formato JSON o in un altro formato comprensibile per il frontend
	nextMoveJSON, err := json.Marshal(nextMove)
	if err != nil {
		return err
	}

	// Invia la griglia di gioco aggiornata a entrambi i giocatori
	if err := player1.WriteMessage(websocket.TextMessage, nextMoveJSON); err != nil {
		return err
	}
	if err := player2.WriteMessage(websocket.TextMessage, nextMoveJSON); err != nil {
		return err
	}

	return nil
}

// UpdateGameStatus aggiorna lo stato della partita e lo invia a entrambi i giocatori.
func UpdateGameStatus(player1, player2 *websocket.Conn, gameStatus GameStatus) error {
	// Converti gameStatus in formato JSON o in un altro formato comprensibile per il frontend
	gameStatusJSON, err := json.Marshal(gameStatus)
	if err != nil {
		return err
	}

	// Invia lo stato della partita aggiornato a entrambi i giocatori
	if err := player1.WriteMessage(websocket.TextMessage, gameStatusJSON); err != nil {
		return err
	}
	if err := player2.WriteMessage(websocket.TextMessage, gameStatusJSON); err != nil {
		return err
	}

	return nil
}
