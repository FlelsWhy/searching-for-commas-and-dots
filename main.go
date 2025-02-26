package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type GameState struct {
	Och    int    `json:"och"`
	Mon    int    `json:"mon"`
	Result string `json:"result"`
}

type Request struct {
	GameType int `json:"gameType"`
	Number   int `json:"number"`
	Color    int `json:"color"`
	Bet      int `json:"bet"`
}

var state = GameState{Och: 20, Mon: 100}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/play", playHandler)
	http.HandleFunc("/state", stateHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}

func stateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(state)
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req Request
	json.NewDecoder(r.Body).Decode(&req)

	switch req.GameType {
	case 1:
		playKosti(req.Number, req.Bet)
	case 2:
		playRuletka(req.Number, req.Color, req.Bet)
	}

	json.NewEncoder(w).Encode(state)
}

func playKosti(number, bet int) {
	winNumber := rand.Intn(6) + 1
	if number == winNumber {
		state.Och += 20
		state.Mon += bet * 3 / 2
		state.Result = fmt.Sprintf("Выигрыш! Выпало %d", winNumber)
	} else {
		state.Och -= 5
		state.Mon -= bet
		state.Result = fmt.Sprintf("Проигрыш! Выпало %d", winNumber)
	}
}

func playRuletka(number, color, bet int) {
	winNumber := rand.Intn(50) + 1
	winColor := rand.Intn(2) + 1

	switch {
	case number == winNumber && color == winColor:
		state.Och += 100
		state.Mon += bet * 50
		state.Result = fmt.Sprintf("Джекпот! Число %d, цвет %d", winNumber, winColor)
	case number == winNumber || color == winColor:
		state.Och += 20
		state.Mon += bet * 2
		state.Result = fmt.Sprintf("Частичный выигрыш! Число %d, цвет %d", winNumber, winColor)
	default:
		state.Och -= 5
		state.Mon -= bet
		state.Result = fmt.Sprintf("Проигрыш! Число %d, цвет %d", winNumber, winColor)
	}
}
