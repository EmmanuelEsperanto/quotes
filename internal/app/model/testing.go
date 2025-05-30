package model

import "testing"

func TestQuote(t *testing.T) *Quote {
	return &Quote{
		Author: "Nicollo Machiavelli",
		Text:   "power tip",
	}
}
