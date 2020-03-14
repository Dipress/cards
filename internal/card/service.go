package card

import (
	"context"
	"fmt"
)

// go:generate mockgen -source=service.go -package=card -destination=service.mock.go

// Repository allows to work with the database.
type Repository interface {
	Create(context.Context, *NewCard, *Card) error
}

// Validater validates card's fields.
type Validater interface {
	Validate(ctx context.Context, form *Form) error
}

// Service is a use case for card creation.
type Service struct {
	Repository
}

// NewService factory prepares service for all futher operations.
func NewService(r Repository) *Service {
	s := Service{
		Repository: r,
	}

	return &s
}

// Create creates a card.
func (s *Service) Create(ctx context.Context, f *Form) (*Card, error) {
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
