/*
Envi (ENron VIsualizer) is a webserver to easily visualize the enron database.
It uses ZincSearch as a backend.

Usage:

	Envi [flags]

Flags:

	-address (optional)
		On what address to run.
		By default, it's ":8080".

	-profiler (optional)
		Whether to have the profiler active or inactive.
		Accepts the values "true" and "false".

	-help
		Brings up a simple explanation of the program flags.
*/
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	addr := flag.String("address", ":8080", "The address to host the visualizer on")
	profiler := flag.Bool("profiler", true, "To have the profiler enabled or disabled")
	flag.Parse()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(cors.Handler)

	if *profiler == true {
		r.Mount("/debug", middleware.Profiler())
	}

	r.Post("/api/search", searchHandler)

	fs := http.FileServer(http.Dir("./frontend/dist"))
	r.Handle("/*", fs)

	http.ListenAndServe(*addr, r)
}

type searchRequest struct {
	SearchTerm string `json:"search_term"`
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	var decoded searchRequest

	err := json.NewDecoder(r.Body).Decode(&decoded)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Requested %s\n", decoded.SearchTerm)
}
