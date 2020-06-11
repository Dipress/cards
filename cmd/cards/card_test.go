package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"testing"

	"github.com/dipress/cards/internal/card"
	"github.com/dipress/cards/internal/storage/postgres"
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
		s := setupServer(lis.Addr().String(), nil, services)
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

func TestFindCard(t *testing.T) {
	t.Log("with prepred server")
	{
		db, teardown := postgresDB(t)
		defer teardown()

		lis, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), caseTimeout)
		defer cancel()

		cardRepo := postgres.NewCardRepository(db)

		nc := card.NewCard{
			Word:          "depict",
			Transcription: "diˈpikt",
			Translation:   "изображать",
			UserID:        2,
		}

		var cd card.Card
		if err := cardRepo.Create(ctx, &nc, &cd); err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		services := setupServices(db)
		s := setupServer(lis.Addr().String(), nil, services)
		go s.Serve(lis)
		defer s.Close()

		t.Log("\ttest:0\tshould find a card.")
		{
			req, err := http.NewRequest(http.MethodGet,
				fmt.Sprintf("http://%s/api/v1/cards/%d", s.Addr, cd.ID), nil)
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

		t.Log("\ttest:1\tshould get a not found error")
		{
			req, err := http.NewRequest(http.MethodGet,
				fmt.Sprintf("http://%s/api/v1/cards/%d", s.Addr, 777), nil)
			req.Header.Set("Content-Type", "application/json")

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if resp.StatusCode != http.StatusNotFound {
				t.Errorf("unexpected status code: %d expected: %d", resp.StatusCode, http.StatusNotFound)
			}
		}
	}
}

func TestUpdateCard(t *testing.T) {
	t.Log("with prepred server")
	{
		db, teardown := postgresDB(t)
		defer teardown()

		lis, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), caseTimeout)
		defer cancel()

		cardRepo := postgres.NewCardRepository(db)

		nc := card.NewCard{
			Word:          "pitfall",
			Transcription: "ˈpitˌfôl",
			Translation:   "ловушка",
			UserID:        3,
		}

		var cd card.Card
		if err := cardRepo.Create(ctx, &nc, &cd); err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		services := setupServices(db)
		s := setupServer(lis.Addr().String(), nil, services)
		go s.Serve(lis)
		defer s.Close()

		t.Log("\ttest:0\tshould update a card.")
		{
			cardStr := `{
				"word": "pitfall", 
				"transcription": "ˈpitˌfôl", 
				"translation": "западня", 
				"user_id": 3
			}`

			req, err := http.NewRequest(http.MethodPut,
				fmt.Sprintf("http://%s/api/v1/cards/%d", s.Addr, cd.ID), strings.NewReader(cardStr))
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
				"word": "pitfall", 
				"transcription": "ˈpitˌfôl", 
				"user_id": 3
			}`

			req, err := http.NewRequest(http.MethodPut,
				fmt.Sprintf("http://%s/api/v1/cards/%d", s.Addr, cd.ID), strings.NewReader(cardStr))
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

		t.Log("\ttest:2\tshould get a not found error")
		{
			cardStr := `{
				"word": "pitfall", 
				"transcription": "ˈpitˌfôl", 
				"translation": "западня", 
				"user_id": 3
			}`

			req, err := http.NewRequest(http.MethodPut,
				fmt.Sprintf("http://%s/api/v1/cards/%d", s.Addr, 154), strings.NewReader(cardStr))
			req.Header.Set("Content-Type", "application/json")

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if resp.StatusCode != http.StatusNotFound {
				t.Errorf("unexpected status code: %d expected: %d", resp.StatusCode, http.StatusNotFound)
			}
		}
	}
}

func TestDeleteCard(t *testing.T) {
	t.Log("with prepred server")
	{
		db, teardown := postgresDB(t)
		defer teardown()

		lis, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), caseTimeout)
		defer cancel()

		cardRepo := postgres.NewCardRepository(db)

		nc := card.NewCard{
			Word:          "own",
			Transcription: "ōn",
			Translation:   "владеть",
			UserID:        4,
		}

		var cd card.Card
		if err := cardRepo.Create(ctx, &nc, &cd); err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		services := setupServices(db)
		s := setupServer(lis.Addr().String(), nil, services)
		go s.Serve(lis)
		defer s.Close()

		t.Log("\ttest:0\tshould delete a card.")
		{
			req, err := http.NewRequest(http.MethodDelete,
				fmt.Sprintf("http://%s/api/v1/cards/%d", s.Addr, cd.ID), nil)
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

		t.Log("\ttest:1\tshould get a not found error")
		{
			req, err := http.NewRequest(http.MethodDelete,
				fmt.Sprintf("http://%s/api/v1/cards/%d", s.Addr, 214), nil)
			req.Header.Set("Content-Type", "application/json")

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if resp.StatusCode != http.StatusNotFound {
				t.Errorf("unexpected status code: %d expected: %d", resp.StatusCode, http.StatusNotFound)
			}
		}
	}
}
