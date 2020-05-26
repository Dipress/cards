package postgres

import (
	"context"
	"database/sql"
	"errors"
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

const createCardQuery = `
	INSERT INTO cards (word, transcription, translation, user_id)
	VALUES ($1, $2, $3, $4)
	RETURNING id, user_id, word, transcription, translation, created_at, updated_at
`

// Create inserts a new card into the database.
func (r *CardRepository) Create(ctx context.Context, f *card.NewCard, ca *card.Card) error {
	if err := r.db.QueryRowContext(ctx, createCardQuery, f.Word, f.Transcription, f.Translation, f.UserID).Scan(
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

const findCardQuery = `SELECT * FROM cards WHERE id = $1`

// Find finds a card by id.
func (r *CardRepository) Find(ctx context.Context, id int) (*card.Card, error) {
	var cd card.Card

	if err := r.db.QueryRowContext(ctx, findCardQuery, id).Scan(
		&cd.ID,
		&cd.UserID,
		&cd.Word,
		&cd.Transcription,
		&cd.Translation,
		&cd.CreatedAt,
		&cd.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, card.ErrNotFound
		}

		return nil, fmt.Errorf("query row scan: %w", err)
	}

	return &cd, nil
}

const updateCardQuery = `
	UPDATE 
		cards 
	SET 
		user_id=:user_id, 
		word=:word,
		transcription=:transcription,
		translation=:translation,
		updated_at=now() 
	WHERE 
		id=:id
	`

// Update updates a card by id
func (r *CardRepository) Update(ctx context.Context, id int, ca *card.Card) error {
	stmt, err := r.db.PrepareNamed(updateCardQuery)
	if err != nil {
		return fmt.Errorf("prepare named: %w", err)
	}
	defer stmt.Close()

	if _, err := stmt.ExecContext(ctx, map[string]interface{}{
		"id":            id,
		"user_id":       ca.UserID,
		"word":          ca.Word,
		"transcription": ca.Transcription,
		"translation":   ca.Translation,
	}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return card.ErrNotFound
		}

		return fmt.Errorf("exec context: %w", err)
	}

	return nil
}

const deleteCardQuery = `DELETE FROM cards WHERE id=:id`

// Delete deletes a card by id
func (r *CardRepository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.PrepareNamed(deleteCardQuery)
	if err != nil {
		return fmt.Errorf("prepare named: %w", err)
	}
	defer stmt.Close()

	if _, err := stmt.ExecContext(ctx, map[string]interface{}{
		"id": id,
	}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return card.ErrNotFound
		}

		return fmt.Errorf("exec context: %w", err)
	}

	return nil
}
