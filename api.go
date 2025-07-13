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

type RateLimitOptions struct {
	Max    int
	Window time.Duration
}

type ApiRoute struct {
	Path             string
	Method           string
	Handler          http.Handler
	Title            string
	Description      string
	RateLimitOptions RateLimitOptions
	Input            any
	Output           any
}

func (app *Api) Mount() {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(60 * time.Second))

	routes := []ApiRoute{
		{
			Path:   "/v1",
			Method: "GET",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("TEST"))
			}),
			Title:            "Test",
			Description:      "Test",
			RateLimitOptions: RateLimitOptions{Max: 20, Window: time.Minute},
			Input:            `{"example": "value"}`,
			Output:           `{"example": "value"}`,
		},
	}

	for _, route := range routes {
		mux.Route(route.Path, func(r chi.Router) {
			r.Method(route.Method, "/", route.Handler)
		})
	}

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
