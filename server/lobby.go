package main

import (
	"sync"
)

// Lobby rappresenta la lobby del gioco, che tiene traccia degli utenti, delle partite e dei giocatori in attesa.
type Lobby struct {
	users   map[string]*User  // Mappa degli utenti con ID come chiave
	matches map[string]*Match // Mappa delle partite con ID come chiave
	waiting []*User           // Slice degli utenti in attesa di una partita
	mu      sync.Mutex        // Mutex per garantire accesso sicuro alla lobby
}

// NewLobby crea una nuova istanza di Lobby.
func NewLobby() *Lobby {
	return &Lobby{
		users:   make(map[string]*User),
		matches: make(map[string]*Match),
	}
}

// AddUser aggiunge un utente alla lobby.
func (l *Lobby) AddUser(user *User) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.users[user.ID] = user
}

// RemoveUser rimuove un utente dalla lobby.
func (l *Lobby) RemoveUser(userID string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	delete(l.users, userID)
}

// AddToWaiting aggiunge un utente alla lista degli utenti in attesa.
func (l *Lobby) AddToWaiting(user *User) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.waiting = append(l.waiting, user)
}

// RemoveFromWaiting rimuove un utente dalla lista degli utenti in attesa.
func (l *Lobby) RemoveFromWaiting(userID string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	for i, user := range l.waiting {
		if user.ID == userID {
			l.waiting = append(l.waiting[:i], l.waiting[i+1:]...)
			break
		}
	}
}

// Match rappresenta una partita in corso tra due giocatori.
type Match struct {
	ID         string           // ID univoco della partita
	Users      map[string]*User // Mappa degli utenti con ID come chiave
	NextPlayer string           // Prossimo giocatore che deve giocare
	Status     GameStatus       // Stato della partita
	Winner     string           // ID del giocatore vincitore
	// TODO da terminare --- non so se è da termianre o meno
}

// CreateMatch crea una nuova partita tra due giocatori.
func (l *Lobby) CreateMatch(player1, player2 *User) *Match {
	// Crea una nuova partita e aggiungila alla lobby
	match := &Match{
		// Inizializza campi della partita
	}
	matchID := generateUniqueMatchID(player1.ID, player2.ID) // Implementa una funzione per generare un ID univoco per la partita
	l.matches[matchID] = match
	return match
}

func generateUniqueMatchID(player1ID, player2ID string) string {
	return player1ID + "." + player2ID
}

// RemoveMatch rimuove una partita dalla lobby.
func (l *Lobby) RemoveMatch(matchID string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	delete(l.matches, matchID)
}