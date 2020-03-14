package card

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dipress/cards/internal/broker/http/handler"
	"github.com/dipress/cards/internal/broker/http/response"
	"github.com/dipress/cards/internal/card"
	"github.com/gorilla/mux"
)

// go:generate mockgen -source=handler.go -package=card -destination=handler.mock.go Service

// Handler allows to handle requests.
type Handler interface {
	Handle(w http.ResponseWriter, r *http.Request) error
}

// Service contains all services.
type Service interface {
	Create(context.Context, *card.Form) (*card.Card, error)
}

// CreateHandler for create requests.
type CreateHandler struct {
	Service
}

// Handle implements Handler interface.
func (h *CreateHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	var f card.Form

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("bad request: %w", response.BadRequest(w))
	}

	if err := f.UnmarshalJSON(data); err != nil {
		return fmt.Errorf("unmarshal json: %w", response.BadRequest(w))
	}

	card, err := h.Create(r.Context(), &f)
	if err != nil {
		return fmt.Errorf("create: %w", response.InternalServerError(w))
	}

	data, err = card.MarshalJSON()
	if err != nil {
		return fmt.Errorf("marshal json: %w", response.InternalServerError(w))
	}

	if _, err := w.Write(data); err != nil {
		return fmt.Errorf("write response: %w", response.InternalServerError(w))
	}

	return nil
}

// Prepare prepares routes to use.
func Prepare(subrouter *mux.Router, service Service, middleware func(handler.Handler) http.Handler) {
	create := CreateHandler{service}

	subrouter.Handle("", middleware(&create)).Methods(http.MethodPost)
}
