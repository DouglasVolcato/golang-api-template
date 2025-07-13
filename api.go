package main

import (
	"api-template/env"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Api struct {
	mux http.Handler
}

func NewApi() *Api {
	return &Api{}
}

func (app *Api) Mount() {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(60 * time.Second))

	mux.Route("/v1", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("TESTE"))
		})
	})

	app.mux = mux
}

func (app *Api) Run() error {
	port := env.GetString("PORT", ":3000")
	svr := &http.Server{
		Addr:         port,
		Handler:      app.mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server started at %s", port)

	return svr.ListenAndServe()
}
