package api

import (
	"api-server/auth"
	"api-server/handlers"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func StartAPI(Port string) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Post("/Login", auth.LoginHandler)
	r.Post("/Book", auth.AuthMiddleWare(handlers.HandleAddBook))
	r.Get("/Book", auth.AuthMiddleWare(handlers.HandleGetBook))
	r.Delete("/Book/{bookId}", auth.AuthMiddleWare(handlers.HandleDeleteBook))
	r.Put("/Book/{bookId}", auth.AuthMiddleWare(handlers.HandleUpdateBook))
	r.Post("/Author", auth.AuthMiddleWare(handlers.HandleAddAuthor))
	r.Get("/Author", auth.AuthMiddleWare(handlers.HandleGetAuthor))
	r.Delete("/Author/{authorId}", auth.AuthMiddleWare(handlers.HandleDeleteAuthor))
	Port = ":" + Port
	fmt.Println(Port)
	http.ListenAndServe(Port, r)
}
