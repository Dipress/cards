package http

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dipress/cards/internal/broker/http/handler"
	"github.com/golang/mock/gomock"
)

func Test_finalizeMiddleware(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := NewMockLogger(ctrl)

	err := errors.New("mock error")

	next := handler.Func(func(w http.ResponseWriter, r *http.Request) error {
		return err
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()

	logger.EXPECT().Error(fmt.Errorf("serve http error: %w", err), map[string]interface{}{
		"path":   "/",
		"method": "GET",
	})

	chain := handler.NewChain(contentTypeMiddleware)

	finalizeMiddleware(logger, chain)(next).ServeHTTP(resp, req)

}
