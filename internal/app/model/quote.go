package model

import (
	"errors"
)

type Quote struct {
	Author string `json:"author"`
	Text   string `json:"quote"`
}

func (q *Quote) Validate() error {
	if q.Author == "" {
		return errors.New("author name is required")
	}
	if len(q.Text) == 0 {
		return errors.New("quote is required")
	}
	return nil
}
