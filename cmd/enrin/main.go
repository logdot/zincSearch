/*
Enrin (ENRon INdexer) is used to automatically feed the enron email database into the ZincSearch service.
It searches the database recursively for all the emails, parses them, and send them to be ingested by ZincSearch.

Usage:

	enrin [flags] [path]

Flags:

	-address (optional)
		The address of the ZincSearch server.
		By default, it's "http://127.0.0.1:4080/".

	-index (required)
		The ZincSearch index to index the database into.

	-password (optional)
		The password for the ZincSearch server.
		By default, it's "admin".

	-profiler (optional)
		The address to run the profiler on.
		It accepts standard ListenAndServe addresses (e.g. ":3000" is valid).
		If left unspecified the profiler will not run.

	-username (optional)
		The username for the ZincSearch server.
		By default, it's "admin".

	-chunking (optional)
		How many files to send to ZincSearch at once.
		Be careful as the larger this number the higher the memory usage.
		Really high numbers might also overwhelm ZincSearch.
		By default, it's 1000.

	-concurrency (optional)
		How many files to index at once.
		Be careful as the larger this number the higher the memory usage.
		By default, it's 500.

	-help
		Brings up a simple explanation of all the program flags.

Path:

	The path to the enron email database.
*/
package main

import (
	"flag"
	"fmt"
	"github.com/logdot/zincSearch"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	username := flag.String("username", "admin", "The username for the ZincSearch server")
	password := flag.String("password", "admin", "The password for the ZincSearch server")
	index := flag.String("index", "", "The ZincSearch index to put the enron emails in")
	address := flag.String("address", "http://127.0.0.1:4080/", "The root address of the ZincSearch server")
	profilerAddr := flag.String("profiler", "", "The address to run the profiling server on. Leave blank to disable")
	concurrency := flag.Uint("concurrency", 500, "How many files to index at once. Beware of memory usage")
	chunking := flag.Uint("chunking", 1000, "How many files to send to ZincSearch at once. Beware of memory and overloading ZincSearch")
	flag.Parse()
	dbPath := flag.Arg(0)

	exit := false
	if *index == "" {
		fmt.Println("Please specify an index to ingest the database into")
		exit = true
	}

	if dbPath == "" {
		fmt.Println("Please specify the path to the enron database")
		exit = true
	}

	if exit {
		return
	}

	if *profilerAddr != "" {
		go http.ListenAndServe(*profilerAddr, nil)
	}

	zincSearch.Authenticate(*address, *index, *username, *password).IndexDB(dbPath, *concurrency, *chunking)
}
