package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Mount("/debug", middleware.Profiler())

	r.Get("/", index)

	http.ListenAndServe(":3000", r)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
