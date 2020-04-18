package card

import (
	"context"
	"fmt"
)

// go:generate mockgen -source=service.go -package=card -destination=service.mock.go

// Repository allows to work with the database.
type Repository interface {
	Create(context.Context, *NewCard, *Card) error
	Find(context.Context, int) (*Card, error)
	Update(context.Context, int, *Card) error
}

// Validater validates card's fields.
type Validater interface {
	Validate(context.Context, *Form) error
}

// Service is a use case for card creation.
type Service struct {
	Repository
	Validater
}

// NewService factory prepares service for all futher operations.
func NewService(r Repository, v Validater) *Service {
	s := Service{
		Repository: r,
		Validater:  v,
	}

	return &s
}

// Create creates a card.
func (s *Service) Create(ctx context.Context, f *Form) (*Card, error) {
	if err := s.Validater.Validate(ctx, f); err != nil {
		return nil, fmt.Errorf("validater validate: %w", err)
	}

	var nc NewCard
	nc.Word = f.Word
	nc.Transcription = f.Transcription
	nc.Translation = f.Translation
	nc.UserID = f.UserID

	var card Card
	if err := s.Repository.Create(ctx, &nc, &card); err != nil {
		return nil, fmt.Errorf("repository create: %w", err)
	}

	return &card, nil

}

// Find finds a card.
func (s *Service) Find(ctx context.Context, id int) (*Card, error) {
	c, err := s.Repository.Find(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("repository find: %w", err)
	}

	return c, nil
}

func (s *Service) Update(ctx context.Context, id int, f *Form) (*Card, error) {
	if err := s.Validater.Validate(ctx, f); err != nil {
		return nil, fmt.Errorf("validater validate: %w", err)
	}

	c, err := s.Repository.Find(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("repository find: %w", err)
	}

	c.UserID = f.UserID
	c.Word = f.Word
	c.Transcription = f.Transcription
	c.Translation = f.Translation

	if err := s.Repository.Update(ctx, id, c); err != nil {
		return nil, fmt.Errorf("repository update: %w", err)
	}

	return c, err
}
