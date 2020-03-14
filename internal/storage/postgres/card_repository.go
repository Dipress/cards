package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dipress/cards/internal/card"
	"github.com/jmoiron/sqlx"
)

const (
	driverName = "postgres"
)

// CardRepository holds CRUD actions.
type CardRepository struct {
	db *sqlx.DB
}

// NewCardRepository factory prepares the card repository to work.
func NewCardRepository(db *sql.DB) *CardRepository {
	r := CardRepository{
		db: sqlx.NewDb(db, driverName),
	}

	return &r
}

const createCard = `
	INSERT INTO cards (word, transcription, translation, user_id)
	VALUES ($1, $2, $3, $4)
	RETURNING id, user_id, word, transcription, translation, created_at, updated_at
`

// Create inserts a new card into the database.
func (r *CardRepository) Create(ctx context.Context, f *card.NewCard, ca *card.Card) error {
	if err := r.db.QueryRowContext(ctx, createCard, f.Word, f.Transcription, f.Translation, f.UserID).Scan(
		&ca.ID,
		&ca.UserID,
		&ca.Word,
		&ca.Transcription,
		&ca.Translation,
		&ca.CreatedAt,
		&ca.UpdatedAt,
	); err != nil {
		return fmt.Errorf("query context scan: %w", err)
	}

	return nil
}
