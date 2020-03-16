package response

import (
	json "encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/dipress/cards/internal/validation"
)

// easyjson -all responses.go

var (
	// ErrBadRequest raises when unable to decode json.
	ErrBadRequest = errors.New("bad request")
)

// HandleError allows to handle default errors.
func HandleError(err error, w http.ResponseWriter) error {
	var vErr validation.Errors
	switch {
	case errors.As(err, &vErr):
		return ValidationError(w, vErr)
	case errors.Is(err, ErrBadRequest):
		return BadRequest(w)
	}

	if rErr := InternalServerError(w); rErr != nil {
		return fmt.Errorf("internal error: %v: %w", rErr, err)
	}

	return fmt.Errorf("internal error: %w", err)
}

// InternalServerError with code 500.
func InternalServerError(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusInternalServerError)
	return writeError(w, "internal server error")
}

// BadRequest responds code 400.
func BadRequest(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusBadRequest)
	return writeError(w, "bad request")
}

type validationResponse struct {
	Message string            `json:"message"`
	Errors  validation.Errors `json:"errors"`
}

// ValidationError responds with code 422.
func ValidationError(w http.ResponseWriter, ers validation.Errors) error {
	w.WriteHeader(http.StatusUnprocessableEntity)

	if err := json.NewEncoder(w).Encode(&ers); err != nil {
		return fmt.Errorf("encode: %w", err)
	}

	return nil
}

type messageResponse struct {
	Message string `json:"message"`
}

func writeError(w http.ResponseWriter, message string) error {
	resp := messageResponse{
		Message: message,
	}

	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		return fmt.Errorf("encode: %w", err)
	}

	return nil
}
