package card

import (
	"errors"
	"time"
)

// easyjson -all model.go

// ErrNotFound raises when role isn't found in the database.
var ErrNotFound = errors.New("card not found")

// constains all card fields.
type Card struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	Word          string    `json:"word"`
	Transcription string    `json:"transcription"`
	Translation   string    `json:"translation"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// NewCard contains the information which needs to create a new Card.
type NewCard struct {
	UserID        int    `json:"user_id"`
	Word          string `json:"word"`
	Transcription string `json:"transcription"`
	Translation   string `json:"translation"`
}

// Form is a card form.
type Form struct {
	UserID        int    `json:"user_id"`
	Word          string `json:"word"`
	Transcription string `json:"transcription"`
	Translation   string `json:"translation"`
}

// Cards contains slice of the cards.
type Cards struct {
	Cards []Card `json:"cards"`
}
