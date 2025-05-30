package apiserver

import (
	"net/http"
	"quotes/internal/app/store/mainstore"
)

func Start() error {
	store := mainstore.New()
	srv := newServer(store)

	return http.ListenAndServe(":8080", srv)
}
