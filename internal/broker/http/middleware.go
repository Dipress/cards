package http

import (
	"net/http"

	"github.com/dipress/cards/internal/broker/http/handler"
)

// contentTypeMiddleware sets content type header.
func contentTypeMiddleware(next handler.Handler) handler.Handler {
	h := handler.Func(func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("Content-Type", "application/json")
		return next.Handle(w, r)
	})

	return h
}
