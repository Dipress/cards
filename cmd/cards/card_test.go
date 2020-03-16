package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"testing"
)

func TestCreateCard(t *testing.T) {
	t.Log("with prepred server")
	{
		db, teardown := postgresDB(t)
		defer teardown()

		lis, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		services := setupServices(db)
		s := setupServer(lis.Addr().String(), services)
		go s.Serve(lis)
		defer s.Close()

		t.Log("\ttest:0\tshould create a card.")
		{
			cardStr := `{
				"word": "do", 
				"transcription": "do͞o", 
				"translation": "делать", 
				"user_id": 1
			}`
			req, err := http.NewRequest(http.MethodPost,
				fmt.Sprintf("http://%s/api/v1/cards", s.Addr), strings.NewReader(cardStr))
			req.Header.Set("Content-Type", "application/json")

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if resp.StatusCode != http.StatusOK {
				t.Errorf("unexpected status code: %d expected: %d", resp.StatusCode, http.StatusOK)
			}
		}
		t.Log("\ttest:1\tshould get a validation error")
		{
			cardStr := `{
				"word": "do", 
				"user_id": 1
			}`
			req, err := http.NewRequest(http.MethodPost,
				fmt.Sprintf("http://%s/api/v1/cards", s.Addr), strings.NewReader(cardStr))
			req.Header.Set("Content-Type", "application/json")

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if resp.StatusCode != http.StatusUnprocessableEntity {
				t.Errorf("unexpected status code: %d expected: %d", resp.StatusCode, http.StatusUnprocessableEntity)
			}
		}
	}
}
