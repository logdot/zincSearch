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
	"flag"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	addr := flag.String("address", ":8080", "The address to host the visualizer on")
	profiler := flag.Bool("profiler", true, "To have the profiler enabled or disabled")
	flag.Parse()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	if *profiler == true {
		r.Mount("/debug", middleware.Profiler())
	}

	r.Get("/", index)

	http.ListenAndServe(*addr, r)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
