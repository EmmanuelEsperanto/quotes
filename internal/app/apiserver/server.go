package apiserver

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"quotes/internal/app/model"
	"quotes/internal/app/store"
	"strconv"
)

type server struct {
	router *mux.Router
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/quotes", s.addQuote()).Methods("POST")
	s.router.HandleFunc("/quotes", s.getAllQuotes()).Methods("GET")
	s.router.HandleFunc("/quotes/random", s.getRandQuote()).Methods("GET")
	s.router.HandleFunc("/quotes/{id}", s.deleteQuote()).Methods("DELETE")

}

func (s *server) addQuote() http.HandlerFunc {
	type request struct {
		Author string `json:"author"`
		Quote  string `json:"quote"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		quote := &model.Quote{Author: req.Author, Text: req.Quote}
		s.store.Quote().Create(quote)
		s.respond(w, r, http.StatusCreated, nil)
	}
}

func (s *server) getAllQuotes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		author := r.URL.Query().Get("author")
		if author != "" {
			// Фильтрация по автору
			quotes, err := s.store.Quote().FindByAuthor(author)
			if err != nil {
				s.error(w, r, http.StatusNoContent, err)
			}
			s.respond(w, r, http.StatusOK, quotes)
		} else {
			// Все цитаты
			quotes, err := s.store.Quote().GetAll()
			if err != nil {
				s.error(w, r, http.StatusNoContent, err)
			}
			s.respond(w, r, http.StatusOK, quotes)
		}

		s.respond(w, r, http.StatusInternalServerError, nil)
	}
}

func (s *server) getRandQuote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		quote, err := s.store.Quote().GetRand()
		if err != nil {
			s.error(w, r, http.StatusNoContent, err)
			return
		}
		s.respond(w, r, http.StatusOK, quote)
	}
}

func (s *server) deleteQuote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)     // получаем map с параметрами пути
		idStr, ok := vars["id"] // вытягиваем "id"
		if !ok {
			s.error(w, r, http.StatusBadRequest, errors.New("missing id parameter"))
			return
		}

		id, err := strconv.Atoi(idStr) // преобразуем в int
		if err != nil {
			s.error(w, r, http.StatusBadRequest, errors.New("invalid id parameter"))
			return
		}

		err = s.store.Quote().Delete(id)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		w.WriteHeader(http.StatusNoContent) // 204 No Content при успешном удалении
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
