package apiserver

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"quotes/internal/app/model"
	"quotes/internal/app/store/teststore"
	"strings"
	"testing"
)

func TestServer_AddQuote(t *testing.T) {
	store := teststore.New()
	srv := newServer(store)

	body := `{"author":"Seneca","quote":"Luck is what happens when preparation meets opportunity."}`

	req := httptest.NewRequest(http.MethodPost, "/quotes", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	srv.router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected status 201 Created, got %d", w.Code)
	}

	quotes, err := store.Quote().FindByAuthor("Seneca")
	if err != nil {
		t.Fatalf("failed to find quote: %v", err)
	}
	if len(quotes) != 1 {
		t.Fatalf("expected 1 quote, got %d", len(quotes))
	}
	if quotes[0].Text != "Luck is what happens when preparation meets opportunity." {
		t.Fatalf("quote text mismatch: got %q", quotes[0].Text)
	}
}

func TestServer_GetAllQuotes(t *testing.T) {
	store := teststore.New()
	srv := newServer(store)

	q1 := model.TestQuote(t)
	q1.Author = "Author 1"
	store.Quote().Create(q1)

	q2 := model.TestQuote(t)
	q2.Author = "Author 2"
	store.Quote().Create(q2)

	req := httptest.NewRequest(http.MethodGet, "/quotes", nil)
	w := httptest.NewRecorder()

	srv.router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var quotes []model.Quote
	if err := json.NewDecoder(w.Body).Decode(&quotes); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(quotes) != 2 {
		t.Fatalf("expected 2 quotes, got %d", len(quotes))
	}
}

func TestServer_GetRandomQuote(t *testing.T) {
	store := teststore.New()
	srv := newServer(store)

	q1 := &model.Quote{Author: "Marcus Aurelius", Text: "You have power over your mind – not outside events."}
	q2 := &model.Quote{Author: "Seneca", Text: "Luck is what happens when preparation meets opportunity."}

	store.Quote().Create(q1)
	store.Quote().Create(q2)

	req := httptest.NewRequest(http.MethodGet, "/quotes/random", nil)
	w := httptest.NewRecorder()

	srv.router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var resp model.Quote
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Text == "" || resp.Author == "" {
		t.Fatal("expected non-empty quote fields")
	}

	valid := (resp.Text == q1.Text && resp.Author == q1.Author) || (resp.Text == q2.Text && resp.Author == q2.Author)
	if !valid {
		t.Fatalf("unexpected quote returned: %+v", resp)
	}
}

func TestGetQuotesByAuthor(t *testing.T) {
	store := teststore.New()
	srv := newServer(store)

	quote := &model.Quote{
		Author: "Albert Einstein",
		Text:   "Imagination is more important than knowledge.",
	}
	store.Quote().Create(quote)

	req := httptest.NewRequest("GET", "/quotes?author=Albert%20Einstein", nil)
	rr := httptest.NewRecorder()

	srv.router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %d", rr.Code)
	}

	var quotes []model.Quote
	if err := json.NewDecoder(rr.Body).Decode(&quotes); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(quotes) != 1 {
		t.Fatalf("expected 1 quote, got %d", len(quotes))
	}

	if quotes[0].Author != quote.Author || quotes[0].Text != quote.Text {
		t.Fatalf("quote mismatch: got %+v", quotes[0])
	}
}
