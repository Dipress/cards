package response

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dipress/cards/internal/validation"
)

func TestBadRequest(t *testing.T) {
	t.Parallel()

	rec := httptest.NewRecorder()
	BadRequest(rec)

	expect := http.StatusBadRequest
	got := rec.Code

	if got != expect {
		t.Errorf("unexpected status code: %d expected: %d", got, expect)
	}

	body, err := ioutil.ReadAll(rec.Body)
	if err != nil {
		t.Errorf("failed to read recorder body: %v", err)
		return
	}

	expectedBody := `{"message":"bad request"}`

	if !strings.Contains(string(body), expectedBody) {
		t.Errorf("unexpected body:\n\t\t%s\nexpected:\n\t\t%s", body, expectedBody)
	}
}

func TestValidationError(t *testing.T) {
	t.Parallel()

	rec := httptest.NewRecorder()

	ves := validation.NewErrors()
	ves.Details["word"] = "cannot be blank"

	ValidationError(rec, ves)

	expect := http.StatusUnprocessableEntity
	got := rec.Code

	if got != expect {
		t.Errorf("unexpected status code: %d expected: %d", got, expect)
	}

	body, err := ioutil.ReadAll(rec.Body)
	if err != nil {
		t.Errorf("failed to read recorder body: %v", err)
		return
	}

	expectedBody := `{"error":"you have validation errors","details":{"word":"cannot be blank"}}`

	if !strings.Contains(string(body), expectedBody) {
		t.Errorf("unexpected body:\n\t\t%s\nexpected:\n\t\t%s", body, expectedBody)
	}
}

func TestInternalServerError(t *testing.T) {
	t.Parallel()

	rec := httptest.NewRecorder()
	InternalServerError(rec)

	expect := http.StatusInternalServerError
	got := rec.Code

	if got != expect {
		t.Errorf("unexpected status code: %d expected: %d", got, expect)
	}

	body, err := ioutil.ReadAll(rec.Body)
	if err != nil {
		t.Errorf("failed to read recorder body: %v", err)
		return
	}

	expectedBody := `{"message":"internal server error"}`

	if !strings.Contains(string(body), expectedBody) {
		t.Errorf("unexpected body:\n\t\t%s\nexpected:\n\t\t%s", body, expectedBody)
	}
}
