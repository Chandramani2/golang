package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


type application struct {
	config config
	// logger
	// db driver
}



// mount -> handlers (controllers)
func (app *application) mount() http.Handler {
	/*
	    - gorilla/mux
		- go-chi/chi
		- fiber
	*/
r := chi.NewRouter()

  // A good base middleware stack
  r.Use(middleware.RequestID)  // Important for Rate-limiting
  r.Use(middleware.RealIP)		// Rate-limiting, Get the real IP from the request, not the proxy
  r.Use(middleware.Logger)    // Logs the start and end of each request with the elapsed processing time
  r.Use(middleware.Recoverer) // Recovers from panics without crashing the server and writes a 500 if there was one.

  // Set a timeout value on the request context (ctx), that will signal
  // through ctx.Done() that the request has timed out and further
  // processing should be stopped.
  r.Use(middleware.Timeout(60 * time.Second))

  r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hi, pinging health check"))
  })

//   // RESTy routes for "articles" resource
//   r.Route("/articles", func(r chi.Router) {
//     r.With(paginate).Get("/", listArticles)                           // GET /articles
//     r.With(paginate).Get("/{month}-{day}-{year}", listArticlesByDate) // GET /articles/01-16-2017

//     r.Post("/", createArticle)                                        // POST /articles
//     r.Get("/search", searchArticles)                                  // GET /articles/search

//     // Regexp url parameters:
//     r.Get("/{articleSlug:[a-z-]+}", getArticleBySlug)                // GET /articles/home-is-toronto

//     // Subrouters:
//     r.Route("/{articleID}", func(r chi.Router) {
//       r.Use(ArticleCtx)
//       r.Get("/", getArticle)                                          // GET /articles/123
//       r.Put("/", updateArticle)                                       // PUT /articles/123
//       r.Delete("/", deleteArticle)                                    // DELETE /articles/123
//     })
//   })

//   // Mount the admin sub-router
//   r.Mount("/admin", adminRouter())

//   http.ListenAndServe(":3333", r)

	return r

	// r.Get("/v1/healthcheck", app.healthcheckHandler)
	// r.Get("/v1/notes", app.notesHandler)
	// r.Get("/v1/notes/{id}", app.noteHandler)

	//http.HandleFunc("/v1/healthcheck", app.healthcheckHandler)
	//http.HandleFunc("/v1/notes", app.notesHandler)
	//http.HandleFunc("/v1/notes/", app.noteHandler)
}

//run  -> graceful shutdown
func(app *application) run(h http.Handler) error{
	// http.ListenAndServe(app.config.addr, app.mount())
	srv := &http.Server{
		Addr: app.config.addr,
		Handler: h,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}
	log.Printf("Server has started at addr %s", app.config.addr)
	return srv.ListenAndServe()

}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string // user,password, dbName, host, port
}