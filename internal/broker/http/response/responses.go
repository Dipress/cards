package response

import (
	"fmt"
	"net/http"
)

// easyjson -all responses.go

// InternalServerError responds code 500.
func InternalServerError(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusInternalServerError)
	return writeError(w, "internal server error")
}

// BadRequest responds code 400.
func BadRequest(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusBadRequest)
	return writeError(w, "bad request")
}

type messageResponse struct {
	Message string `json:"message"`
}

func writeError(w http.ResponseWriter, message string) error {
	resp := messageResponse{
		Message: message,
	}

	data, err := resp.MarshalJSON()
	if err != nil {
		return fmt.Errorf("marshal json: %w", err)
	}

	if _, err := w.Write(data); err != nil {
		return fmt.Errorf("write response: %w", err)
	}
	return nil
}
