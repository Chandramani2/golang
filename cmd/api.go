package main

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


type application struct {
	config config
	// logger
	// db driver
}

//run  -> graceful shutdown

// mount -> handlers (controllers)
func (app *application) mount() http.Handler {
	/*
	    - gorilla/mux
		- go-chi/chi
		- fiber
	*/

	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})

	// http.ListenAndServe(":3333", r)

	return r

	// r.Get("/v1/healthcheck", app.healthcheckHandler)
	// r.Get("/v1/notes", app.notesHandler)
	// r.Get("/v1/notes/{id}", app.noteHandler)

	//http.HandleFunc("/v1/healthcheck", app.healthcheckHandler)
	//http.HandleFunc("/v1/notes", app.notesHandler)
	//http.HandleFunc("/v1/notes/", app.noteHandler)
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string // user,password, dbName, host, port
}