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
	r.Post("/Login", LoginHandler)
	r.Post("/Book", AuthMiddleWare(HandleAddBook))
	r.Get("/Book", AuthMiddleWare(HandleGetBook))
	r.Delete("/Book/{bookId}", AuthMiddleWare(HandleDeleteBook))
	r.Put("/Book/{bookId}", AuthMiddleWare(HandleUpdateBook))
	r.Post("/Author", AuthMiddleWare(HandleAddAuthor))
	r.Get("/Author", AuthMiddleWare(HandleGetAuthor))
	r.Delete("/Author/{authorId}", AuthMiddleWare(HandleDeleteAuthor))
	http.ListenAndServe(":8081", r)
}
