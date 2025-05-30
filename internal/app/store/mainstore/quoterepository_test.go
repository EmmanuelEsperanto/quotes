package mainstore_test

import (
	"quotes/internal/app/model"
	"quotes/internal/app/store"
	"quotes/internal/app/store/teststore"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	q := model.TestQuote(t)

	if err := s.Quote().Create(q); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if q == nil {
		t.Fatal("expected quote to be not nil")
	}
}

func TestUserRepository_FindByAuthor(t *testing.T) {
	s := teststore.New()
	author := "Niccolo Machiavelli"

	_, err := s.Quote().FindByAuthor(author)
	if err == nil || err.Error() != store.ErrRecordNotFound.Error() {
		t.Fatalf("expected error '%v', got '%v'", store.ErrRecordNotFound, err)
	}

	q := model.TestQuote(t)
	q.Author = author

	if err := s.Quote().Create(q); err != nil {
		t.Fatalf("unexpected error on create: %v", err)
	}

	quotes, err := s.Quote().FindByAuthor(author)
	if err != nil {
		t.Fatalf("unexpected error on find: %v", err)
	}
	if quotes == nil {
		t.Fatal("expected quotes to be not nil")
	}
}

func TestUserRepository_Delete(t *testing.T) {
	s := teststore.New()
	q1 := model.TestQuote(t)

	if err := s.Quote().Create(q1); err != nil {
		t.Fatalf("unexpected error on create: %v", err)
	}

	if err := s.Quote().Delete(0); err != nil {
		t.Fatalf("unexpected error on delete: %v", err)
	}
}
