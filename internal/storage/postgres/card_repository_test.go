package postgres

import (
	"context"
	"testing"

	"github.com/dipress/cards/internal/card"
)

func TestCreateCard(t *testing.T) {
	t.Log("with initialized repository")
	{
		db, teardown := postgresDB(t)
		defer teardown()

		r := NewCardRepository(db)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		t.Log("\ttest:0\tshould create the card into the database")
		{
			nc := card.NewCard{
				UserID:        1,
				Word:          "exceed",
				Transcription: "ikˈsēd",
				Translation:   "превышать",
			}

			var cd card.Card
			err := r.Create(ctx, &nc, &cd)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if cd.ID == 0 {
				t.Error("expected to parse returned id")
			}
		}
	}
}

func TestFindCard(t *testing.T) {
	t.Log("with initialized repository")
	{
		db, teardown := postgresDB(t)
		defer teardown()

		r := NewCardRepository(db)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		nc := card.NewCard{
			UserID:        2,
			Word:          "exceed",
			Transcription: "ikˈsēd",
			Translation:   "превышать",
		}

		var cd card.Card
		err := r.Create(ctx, &nc, &cd)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		t.Log("\ttest:0\tshould find the card into the database")
		{
			_, err := r.Find(ctx, cd.ID)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		}
	}
}

func TestUpdateCard(t *testing.T) {
	t.Log("with initialized repository")
	{
		db, teardown := postgresDB(t)
		defer teardown()

		r := NewCardRepository(db)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		nc := card.NewCard{
			UserID:        3,
			Word:          "grow",
			Transcription: "grō",
			Translation:   "расти",
		}

		var cd card.Card
		err := r.Create(ctx, &nc, &cd)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		t.Log("\ttest:0\tshould update the card into the database")
		{
			cd.Word = "climb"
			cd.Transcription = "klīm"
			cd.Translation = "взбираться"

			err := r.Update(ctx, 3, &cd)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		}
	}
}

func TestDeleteCard(t *testing.T) {
	t.Log("with initialized repository")
	{
		db, teardown := postgresDB(t)
		defer teardown()

		r := NewCardRepository(db)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		nc := card.NewCard{
			UserID:        4,
			Word:          "srpead",
			Transcription: "spred",
			Translation:   "распространять",
		}

		var cd card.Card
		err := r.Create(ctx, &nc, &cd)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		t.Log("\ttest:0\tshould delete the card into the database")
		{
			err := r.Delete(ctx, cd.ID)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

		}
	}
}
