package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dipress/cards/internal/broker/http/handler"
)

func Test_contentTypeMIddleware(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "http://exapmle.com", nil)
	rec := httptest.NewRecorder()

	next := handler.Func(func(w http.ResponseWriter, r *http.Request) error {
		return nil
	})

	contentTypeMiddleware(next).Handle(rec, req)

	ct := rec.Header().Get("Content-Type")
	if ct != "application/json" {
		t.Errorf("expected to set application/json Content-Type header: %s", ct)
	}
}
