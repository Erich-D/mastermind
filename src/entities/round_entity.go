package entities

import "time"

type Round struct {
	ID          string    `json:"id"`
	Guess       string    `json:"guess"`
	GuessTime   time.Time `json:"guessTime"`
	GuessResult string    `json:"guessresult"`
	GameID      string    `json:"gameId"`
}
