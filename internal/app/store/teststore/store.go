package teststore

import (
	"quotes/internal/app/model"
	"quotes/internal/app/store"
)

type Store struct {
	quoteRepository *QuoteRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) Quote() store.QuoteRepository {
	if s.quoteRepository != nil {
		return s.quoteRepository
	}

	s.quoteRepository = &QuoteRepository{
		store:       s,
		Quotes:      make(map[int]*model.Quote),
		nextQuoteID: -1,
	}
	return s.quoteRepository
}
