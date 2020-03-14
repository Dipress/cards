package validation

import (
	"context"
	"reflect"
	"testing"

	"github.com/dipress/cards/internal/card"
)

func TestCardValidate(t *testing.T) {
	tests := []struct {
		name    string
		form    card.Form
		wantErr bool
		expect  Errors
	}{
		{
			name: "ok",
			form: card.Form{
				Word:          "make",
				Transcription: "māk",
				Translation:   "сделать",
			},
		},
		{
			name: "blank word",
			form: card.Form{
				Transcription: "māk",
				Translation:   "сделать",
			},
			wantErr: true,
			expect:  Errors{"word": "cannot be blank"},
		},
		{
			name: "blank transcription",
			form: card.Form{
				Word:        "make",
				Translation: "сделать",
			},
			wantErr: true,
			expect:  Errors{"transcription": "cannot be blank"},
		},
		{
			name: "blank translation",
			form: card.Form{
				Word:          "make",
				Transcription: "māk",
			},
			wantErr: true,
			expect:  Errors{"translation": "cannot be blank"},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			var c Card
			err := c.Validate(ctx, &tc.form)
			if tc.wantErr {
				got, ok := err.(Errors)
				if !ok {
					t.Errorf("unknown error: %v", err)
					return
				}

				if !reflect.DeepEqual(tc.expect, got) {
					t.Errorf("expected: %+#v got: %+#v", tc.expect, got)
				}

				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
