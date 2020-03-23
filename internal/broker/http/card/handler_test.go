package card

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dipress/cards/internal/card"
	"github.com/dipress/cards/internal/validation"
	gomock "github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

func TestCreateHandler(t *testing.T) {
	tests := []struct {
		name        string
		serviceFunc func(mock *MockService)
		code        int
	}{
		{
			name: "ok",
			serviceFunc: func(m *MockService) {
				m.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&card.Card{}, nil)
			},
			code: http.StatusOK,
		},
		{
			name: "validation",
			serviceFunc: func(m *MockService) {
				var ves validation.Errors
				m.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, ves)
			},
			code: http.StatusUnprocessableEntity,
		},
		{
			name: "internal error",
			serviceFunc: func(m *MockService) {
				m.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&card.Card{}, errors.New("mock error"))
			},
			code: http.StatusInternalServerError,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := NewMockService(ctrl)
			tc.serviceFunc(service)

			h := CreateHandler{service}
			w := httptest.NewRecorder()

			r := httptest.NewRequest(http.MethodPost, "http://example.com", strings.NewReader("{}"))

			err := h.Handle(w, r)
			if w.Code != tc.code {
				t.Errorf("unexpected code: %d expected %d error: %v", w.Code, tc.code, err)
			}
		})
	}
}

func TestFindHandler(t *testing.T) {
	tests := []struct {
		name        string
		serviceFunc func(mock *MockService)
		code        int
	}{
		{
			name: "ok",
			serviceFunc: func(m *MockService) {
				m.EXPECT().Find(gomock.Any(), gomock.Any()).Return(&card.Card{}, nil)
			},
			code: http.StatusOK,
		},
		{
			name: "not found error",
			serviceFunc: func(m *MockService) {
				m.EXPECT().Find(gomock.Any(), gomock.Any()).Return(&card.Card{}, card.ErrNotFound)
			},
			code: http.StatusNotFound,
		},
		{
			name: "internal error",
			serviceFunc: func(m *MockService) {
				m.EXPECT().Find(gomock.Any(), gomock.Any()).Return(&card.Card{}, errors.New("mock error"))
			},
			code: http.StatusInternalServerError,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := NewMockService(ctrl)
			tc.serviceFunc(service)

			h := FindHandler{service}
			w := httptest.NewRecorder()

			r := httptest.NewRequest(http.MethodGet, "http://example.com", strings.NewReader("{}"))
			r = mux.SetURLVars(r, map[string]string{"id": "1"})

			err := h.Handle(w, r)
			if w.Code != tc.code {
				t.Errorf("unexpected code: %d expected %d error: %v", w.Code, tc.code, err)
			}
		})
	}
}
