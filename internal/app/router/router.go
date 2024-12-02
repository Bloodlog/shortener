package router

import (
	"net/http"
	"shortener/internal/app/handlers"
	"shortener/internal/app/repository"
)

func Run(repository *repository.URLRepository) error {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, handlers.Handle(repository))

	return http.ListenAndServe(`:8080`, mux)
}
