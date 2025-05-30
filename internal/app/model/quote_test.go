package model_test

import (
	"quotes/internal/app/model"
	"testing"
)

func TestQuote_Validate(t *testing.T) {

	testCases := []struct {
		name    string
		a       func() *model.Quote
		isValid bool
	}{
		{
			name: "Valid",
			a: func() *model.Quote {
				return model.TestQuote(t)
			},
			isValid: true,
		},
		{
			name: "Empty Author",
			a: func() *model.Quote {
				q := model.TestQuote(t)
				q.Author = ""
				return q
			},
			isValid: false,
		},
		{
			name: "empty quote text",
			a: func() *model.Quote {
				q := model.TestQuote(t)
				q.Text = ""
				return q
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				err := tc.a().Validate()
				if err != nil {
					t.Errorf("Quote.Validate() error = %v", err)
				}
			} else {
				err := tc.a().Validate()
				if err == nil {
					t.Error("Quote.Validate() error = we want error but didn't get it", nil)
				}
			}
		})
	}
}
