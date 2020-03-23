package card

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
	Create(ctx context.Context, f *card.Form) (*card.Card, error)
	Find(ctx context.Context, id int) (*card.Card, error)
}

// CreateHandler for create requests.
type CreateHandler struct {
	Service
}

// Handle implements Handler interface.
func (h *CreateHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	if err := h.process(w, r); err != nil {
		return response.HandleError(err, w)
	}

	return nil
}

func (h *CreateHandler) process(w http.ResponseWriter, r *http.Request) error {
	var f card.Form

	if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
		return response.ErrBadRequest
	}

	card, err := h.Create(r.Context(), &f)
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}

	if err := json.NewEncoder(w).Encode(&card); err != nil {
		return fmt.Errorf("encode: %w", err)
	}

	return nil
}

// FindHandler for find requests.
type FindHandler struct {
	Service
}

func (h *FindHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	if err := h.process(w, r); err != nil {
		return response.HandleError(err, w)
	}

	return nil
}

func (h *FindHandler) process(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return response.ErrBadRequest
	}

	card, err := h.Service.Find(r.Context(), id)
	if err != nil {
		return fmt.Errorf("find: %w", err)
	}

	if err := json.NewEncoder(w).Encode(&card); err != nil {
		return fmt.Errorf("encode: %w", err)
	}

	return nil
}

// Prepare prepares routes to use.
func Prepare(subrouter *mux.Router, service Service, middleware func(handler.Handler) http.Handler) {
	create := CreateHandler{service}
	find := FindHandler{service}

	subrouter.Handle("", middleware(&create)).Methods(http.MethodPost)
	subrouter.Handle("/{id}", middleware(&find)).Methods(http.MethodGet)
}
