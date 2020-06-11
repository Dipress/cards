package http

import (
	"net/http"
	"time"

	cardHandlers "github.com/dipress/cards/internal/broker/http/card"
	"github.com/dipress/cards/internal/broker/http/handler"
	"github.com/dipress/cards/internal/card"
	"github.com/dipress/cards/internal/kit/logger"
	"github.com/gorilla/mux"
)

const (
	timeout = 15 * time.Second
)

// Services contains all the services.
type Services struct {
	Card *card.Service
}

// NewServer prepares the http server to work.
func NewServer(addr string, logger *logger.Logger, services *Services) *http.Server {
	mux := mux.NewRouter().StrictSlash(true)

	base := handler.NewChain(contentTypeMiddleware)

	cards := mux.PathPrefix("/api/v1/cards").Subrouter()
	cardHandlers.Prepare(cards, services.Card, finalizeMiddleware(logger, base))

	s := http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}

	return &s
}
