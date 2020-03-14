package http

import (
	"log"
	"net/http"
	"time"

	cardHandlers "github.com/dipress/cards/internal/broker/http/card"
	"github.com/dipress/cards/internal/broker/http/handler"
	"github.com/dipress/cards/internal/card"
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
func NewServer(addr string, services *Services) *http.Server {
	mux := mux.NewRouter().StrictSlash(true)

	base := handler.NewChain(contentTypeMiddleware)

	cards := mux.PathPrefix("/api/v1/cards").Subrouter()
	cardHandlers.Prepare(cards, services.Card, finalizeMiddleware(base))

	s := http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}

	return &s
}

func finalizeMiddleware(middleware handler.Chain) func(handler.Handler) http.Handler {
	f := func(handler handler.Handler) http.Handler {
		wrapped := middleware.Then(handler)
		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if err := wrapped.Handle(w, r); err != nil {
				log.Printf("serve http: %+v\n", err)
			}
		})

		return h
	}

	return f
}
