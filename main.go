package main

import (
	"net/http"

	"github.com/FunkyTracy/party-api/partyHandler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	r.Post("/parties", partyHandler.HandleGetParties)

	http.ListenAndServe(":3020", r)
}
