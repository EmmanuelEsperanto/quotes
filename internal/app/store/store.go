package store

type Store interface {
	Quote() QuoteRepository
}
