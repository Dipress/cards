package validation

import (
	"context"

	"github.com/dipress/cards/internal/card"
	validation "github.com/go-ozzo/ozzo-validation"
)

const (
	validationMsg = "you have validation errors"
)

// Errors holds validation errors.
type Errors struct {
	Message string            `json:"error"`
	Details map[string]string `json:"details"`
}

// NewErrors returns prepared errors.
func NewErrors() Errors {
	e := Errors{
		Message: "you have validation errors",
		Details: make(map[string]string),
	}

	return e
}

// Error implements error interface.
func (v Errors) Error() string {
	return v.Message
}

// Card holds form validations.
type Card struct{}

// Validate validates card form.
func (c *Card) Validate(ctx context.Context, form *card.Form) error {
	ves := NewErrors()
	if err := validation.Validate(
		form.Word,
		validation.Required,
	); err != nil {
		ves.Details["word"] = err.Error()
	}

	if err := validation.Validate(
		form.Transcription,
		validation.Required,
	); err != nil {
		ves.Details["transcription"] = err.Error()
	}

	if err := validation.Validate(
		form.Translation,
		validation.Required,
	); err != nil {
		ves.Details["translation"] = err.Error()
	}

	if len(ves.Details) > 0 {
		return ves
	}

	return nil
}
