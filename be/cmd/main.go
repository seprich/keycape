package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/samber/lo"
	"github.com/seprich/keycape/internal/api"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	//r.Use(middleware.ContentCharset())

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("{\"result\":\"ok\"}"))
	})

	r.Route("/api", api.ApiRoutes)

	lo.Must0(http.ListenAndServe(":3000", r))
}
