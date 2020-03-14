package validation

import (
	"context"

	"github.com/dipress/cards/internal/card"
	validation "github.com/go-ozzo/ozzo-validation"
)

const (
	mismatchMsg   = "mismatch"
	validationMsg = "you have validation errors"
)

// Errors holds validation errors.
type Errors map[string]string

// Error implements error interface.
func (v Errors) Error() string {
	return validationMsg
}

// Card holds form validations.
type Card struct{}

// Validate validates card form.
func (c *Card) Validate(ctx context.Context, form *card.Form) error {
	ves := make(Errors)

	if err := validation.Validate(
		form.Word,
		validation.Required,
	); err != nil {
		ves["word"] = err.Error()
	}

	if err := validation.Validate(
		form.Transcription,
		validation.Required,
	); err != nil {
		ves["transcription"] = err.Error()
	}

	if err := validation.Validate(
		form.Translation,
		validation.Required,
	); err != nil {
		ves["translation"] = err.Error()
	}

	if len(ves) > 0 {
		return ves
	}

	return nil
}
