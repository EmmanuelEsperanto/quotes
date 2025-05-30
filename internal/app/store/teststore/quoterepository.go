package teststore

import (
	"errors"
	"math/rand/v2"
	"quotes/internal/app/model"
	"quotes/internal/app/store"
	"strings"
)

type QuoteRepository struct {
	store       *Store
	Quotes      map[int]*model.Quote
	nextQuoteID int
}

func (r *QuoteRepository) Create(q *model.Quote) error {
	if err := q.Validate(); err != nil {
		return err
	}
	r.nextQuoteID++
	r.Quotes[r.nextQuoteID] = q
	return nil
}

func (r *QuoteRepository) GetAll() ([]*model.Quote, error) {
	if len(r.Quotes) == 0 {
		return nil, errors.New("no have quotes for get all")
	}

	quotes := make([]*model.Quote, 0, len(r.Quotes))
	for _, q := range r.Quotes {
		quotes = append(quotes, q)
	}

	return quotes, nil
}

func (r *QuoteRepository) GetRand() (*model.Quote, error) {
	if len(r.Quotes) == 0 {
		return nil, errors.New("no have quotes for rand get")
	}
	randomID := rand.IntN(len(r.Quotes) - 1)
	randomQuote := r.Quotes[randomID]
	return randomQuote, nil
}

func (r *QuoteRepository) FindByAuthor(author string) ([]*model.Quote, error) {
	var result []*model.Quote
	for _, q := range r.Quotes {
		if strings.EqualFold(q.Author, author) {
			result = append(result, q)
		}
	}
	if len(result) == 0 {
		return nil, store.ErrRecordNotFound
	}
	return result, nil
}

func (r *QuoteRepository) Delete(id int) error {
	if len(r.Quotes) == 0 {
		return errors.New("no have quotes for delete")
	}
	for k, _ := range r.Quotes {
		if k == id {
			delete(r.Quotes, id)
			return nil
		}
	}

	return errors.New("cant't find quote")
}
