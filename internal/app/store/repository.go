package store

import "quotes/internal/app/model"

type QuoteRepository interface {
	Create(q *model.Quote) error
	GetAll() ([]*model.Quote, error)
	GetRand() (*model.Quote, error)
	Delete(id int) error
	FindByAuthor(author string) ([]*model.Quote, error)
}
