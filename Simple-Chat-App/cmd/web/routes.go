package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"websocketPractice/internal/handlers"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Home)
	mux.Get("/ws", handlers.WsEndpoint)
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
