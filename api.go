package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Post("/Book", HandleAddBook)
	r.Get("/Book", HandleGetBook)
	r.Delete("/Book/{bookId}", HandleDeleteBook)
	r.Put("/Book/{bookId}", HandleUpdateBook)
	r.Post("/Author", HandleAddAuthor)
	r.Get("/Author", HandleGetAuthor)
	r.Delete("/Author/{authorId}", HandleDeleteAuthor)
	http.ListenAndServe(":8081", r)
}
